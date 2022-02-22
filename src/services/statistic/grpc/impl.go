package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/zikwall/blogchain/src/pkg/clickhouse"
	stdout "github.com/zikwall/blogchain/src/pkg/log"
	"github.com/zikwall/blogchain/src/pkg/maxmind"
	"github.com/zikwall/blogchain/src/protobuf/common"
	"github.com/zikwall/blogchain/src/protobuf/storage"
	"github.com/zikwall/blogchain/src/services/api/utils"
	"github.com/zikwall/blogchain/src/services/statistic/repository"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/mssola/user_agent"
	clickhousebuffer "github.com/zikwall/clickhouse-buffer"
	"github.com/zikwall/clickhouse-buffer/src/buffer/memory"
)

type grpcServerImpl struct {
	storage.UnimplementedStorageServer
	clickhouse      *clickhouse.Connection
	postStatsWriter clickhousebuffer.Writer
	GeoReader       *maxmind.Reader
}

func New(buffer clickhousebuffer.Client, ch *clickhouse.Connection) storage.StorageServer {
	s := &grpcServerImpl{
		clickhouse: ch,
	}
	s.initWriterAPI(buffer)

	return s
}

func (s *grpcServerImpl) initWriterAPI(buffer clickhousebuffer.Client) {
	s.postStatsWriter = buffer.Writer(
		clickhousebuffer.View{
			Name:    repository.PostStatsTable,
			Columns: repository.PostStatsColumns,
		},
		memory.NewBuffer(
			buffer.Options().BatchSize(),
		),
	)
	postStatsErrors := s.postStatsWriter.Errors()
	go func() {
		for err := range postStatsErrors {
			stdout.Warningf("[POST STATS] clickhouse write error: %s\n", err.Error())
		}
	}()
}

func (s *grpcServerImpl) WritePostStats(_ context.Context, in *storage.PostStats) (*common.EmptyResponse, error) {
	now := time.Now()
	stats := &repository.PostStats{
		PostID:   in.PostID,
		OwnerID:  in.OwnerID,
		IP:       in.Ip,
		InsertTS: utils.Datetime(now),
		Date:     utils.Date(now),
	}
	// get user location information
	geo, err := s.GeoReader.Lookup(in.Ip)
	if err == nil {
		stats.Region = geo.Region
		stats.Country = geo.Country
	}
	// get information from user agent
	if in.UserAgent != "" {
		ua := user_agent.New(in.UserAgent)
		stats.Os = ua.OS()
		browser, version := ua.Browser()
		if browser != "" {
			stats.Browser = fmt.Sprintf("%s/%s", browser, version)
		}
		stats.Platform = ua.Platform()
	}
	s.postStatsWriter.WriteRow(stats)
	return nil, nil
}

func (s *grpcServerImpl) GetPostViewersCount(
	ctx context.Context,
	in *storage.PostViewersRequest,
) (
	*storage.PostViewersResponse,
	error,
) {
	var count uint64
	var postID uint64

	rawQuery, _, _ := viewersAggregateQuery().Where(builder.And(builder.C("post_id").Eq(in.PostID))).ToSQL()
	if err := s.clickhouse.Query().QueryRowContext(ctx, rawQuery).Scan(&postID, &count); err != nil {
		return nil, err
	}

	// current
	count++
	return &storage.PostViewersResponse{
		PostID: postID,
		Views:  count,
	}, nil
}

func (s *grpcServerImpl) GetPostsViewersCount(
	ctx context.Context,
	in *storage.PostsViewersRequest,
) (
	*storage.PostsViewersResponse,
	error,
) {
	var views []Viewers
	rawQuery, _, _ := viewersAggregateQuery().Where(builder.And(builder.C("post_id").In(in.PostID))).ToSQL()
	if err := s.clickhouse.Query().SelectContext(ctx, &views, rawQuery); err != nil {
		return nil, err
	}

	response := &storage.PostsViewersResponse{
		Views: make([]*storage.PostViewersResponse, 0, len(views)),
	}
	for _, view := range views {
		response.Views = append(response.Views, &storage.PostViewersResponse{
			PostID: view.PostID,
			Views:  view.Views,
		})
	}
	return response, nil
}

type Viewers struct {
	PostID uint64 `db:"post_id"`
	Views  uint64 `db:"views"`
}

func viewersAggregateQuery() *builder.SelectDataset {
	return builder.
		Select(
			builder.C("post_id"),
			builder.L("sum(views) as views"),
		).
		From("post_stats_views").
		GroupBy(
			builder.C("post_id"),
		)
}
