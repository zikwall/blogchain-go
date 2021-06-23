package actions

import (
	"context"
	"github.com/zikwall/blogchain/src/app/lib/upload"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/maxmind"
	"github.com/zikwall/clickhouse-buffer/src/api"
	"github.com/zikwall/clickhouse-buffer/src/buffer/memory"
)

// The HttpController structure is the base object for all http handlers,
// and encapsulates access to services such as databases, redis, etc.
type HttpController struct {
	RSA              container.RSA
	Db               *database.Connection
	Clickhouse       *clickhouse.Clickhouse
	ClickhouseBuffer *clickhouse.BufferAdapter
	writeAPI         api.Writer
	Finder           *maxmind.Finder
	Uploader         upload.Uploader
}

func CreateHttpControllerWithCopy(context context.Context, p HttpController) *HttpController {
	tableView := api.View{
		Name: "blogchain.post_stats",
		Columns: []string{
			"post_id", "owner_id", "os", "browser", "platform",
			"ip", "country", "region", "insert_ts", "date",
		},
	}

	writeAPI := p.ClickhouseBuffer.Client().Writer(tableView, memory.NewBuffer(
		p.ClickhouseBuffer.Client().Options().BatchSize(),
	))

	return &HttpController{
		RSA:              p.RSA,
		Db:               p.Db,
		Clickhouse:       p.Clickhouse,
		ClickhouseBuffer: p.ClickhouseBuffer,
		Finder:           p.Finder,
		Uploader:         p.Uploader,
		writeAPI:         writeAPI,
	}
}

func (hc HttpController) response(response interface{}) Response {
	return Response{
		Response: response,
	}
}

func (hc HttpController) message(message string) MessageResponse {
	return MessageResponse{
		Status:  200,
		Message: message,
	}
}
