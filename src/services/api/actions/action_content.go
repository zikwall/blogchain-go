package actions

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/pkg/log"
	"github.com/zikwall/blogchain/src/protobuf/storage"
	"github.com/zikwall/blogchain/src/services/api/repositories"

	"github.com/gofiber/fiber/v2"
)

type ContentResponse struct {
	Content repositories.PublicContent `json:"content"`
	Viewers uint64                     `json:"viewers"`
}

type ContentsResponse struct {
	Contents []repositories.PublicContent `json:"contents"`
	Meta     Meta                         `json:"meta"`
	Stats    map[uint64]uint64            `json:"stats"`
}

type Meta struct {
	Pages float64 `json:"pages"`
}

func (hc *HTTPController) Content(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return exceptions.Wrap("failed parse content id", exceptions.ThrowPublicError(err))
	}

	result, err := repositories.UseContentRepository(ctx.Context(), hc.DB).FindContentByID(id)
	if err != nil {
		return exceptions.Wrap("failed find content by id", err)
	}

	defer func() {
		writeContext, cancel := context.WithTimeout(ctx.Context(), time.Second*1)
		defer cancel()
		_, writeErr := hc.StatisticClient.WritePostStats(writeContext, &storage.PostStats{
			PostID:    uint64(result.ID),
			OwnerID:   uint64(result.UserID),
			Ip:        fmt.Sprintf("%v", ctx.Locals("ip")),
			UserAgent: ctx.Get("User-Agent", ""),
		})
		if writeErr != nil {
			log.Warning(err)
		}
	}()

	viewers, err := hc.StatisticClient.GetPostViewersCount(ctx.Context(), &storage.PostViewersRequest{
		PostID: uint64(result.ID),
	})
	if err != nil {
		log.Warning(err)
	}
	return ctx.Status(200).JSON(hc.response(ContentResponse{
		Content: result.GetPublicContent(),
		Viewers: viewers.Views,
	}))
}

func (hc *HTTPController) Contents(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")

	contents, count, err := repositories.UseContentRepository(ctx.Context(), hc.DB).
		FindAllContent(tag, extractPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find contents", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: hc.withStatsContext(ctx.Context(), contents),
	}))
}

func (hc *HTTPController) ContentsUser(ctx *fiber.Ctx) error {
	user, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return exceptions.Wrap("failed parse user id", err)
	}

	contents, count, err := repositories.UseContentRepository(ctx.Context(), hc.DB).
		FindAllByUser(user, extractPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find user contents by id", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: hc.withStatsContext(ctx.Context(), contents),
	}))
}

func (hc *HTTPController) withStatsContext(
	ctx context.Context,
	cs []repositories.PublicContent,
) map[uint64]uint64 {
	if len(cs) == 0 {
		return map[uint64]uint64{}
	}
	ids := make([]uint64, 0, len(cs))
	for i := range cs {
		ids = append(ids, uint64(cs[i].ID))
	}
	viewers, err := hc.StatisticClient.GetPostsViewersCount(ctx, &storage.PostsViewersRequest{
		PostID: ids,
	})
	if err != nil {
		log.Warning(err)
	}
	views := make(map[uint64]uint64, len(viewers.Views))
	for i := range viewers.Views {
		views[viewers.Views[i].PostID] = viewers.Views[i].Views
	}
	return views
}
