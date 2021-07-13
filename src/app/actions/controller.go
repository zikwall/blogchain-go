package actions

import (
	"github.com/zikwall/blogchain/src/app/lib/upload"
	"github.com/zikwall/blogchain/src/app/statistic"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/maxmind"
	"github.com/zikwall/clickhouse-buffer/src/api"
	"github.com/zikwall/clickhouse-buffer/src/buffer/memory"
	"github.com/zikwall/fsclient"
)

// The HTTPController structure is the base object for all http handlers,
// and encapsulates access to services such as databases, redis, etc.
type HTTPController struct {
	RSA              container.RSA
	DB               *database.Connection
	Clickhouse       *clickhouse.Connection
	ClickhouseBuffer *clickhouse.BufferAdapter
	writeAPI         api.Writer
	Finder           *maxmind.Finder
	Uploader         upload.Uploader
	FsClient         *fsclient.FsClient
}

func CreateHTTPControllerWithCopy(p *HTTPController) (*HTTPController, error) {
	controller := &HTTPController{
		RSA:              p.RSA,
		DB:               p.DB,
		Clickhouse:       p.Clickhouse,
		ClickhouseBuffer: p.ClickhouseBuffer,
		Finder:           p.Finder,
		FsClient:         p.FsClient,
	}

	if err := controller.after(); err != nil {
		return nil, err
	}

	return controller, nil
}

func (hc *HTTPController) after() error {
	hc.initWriterAPI()
	return hc.initFileServerClient()
}

func (hc *HTTPController) initWriterAPI() {
	buffer := hc.ClickhouseBuffer.Client()
	hc.writeAPI = buffer.Writer(
		api.View{
			Name:    statistic.PostStatsTable,
			Columns: statistic.PostStatsColumns,
		},
		memory.NewBuffer(
			buffer.Options().BatchSize(),
		),
	)
}

func (hc *HTTPController) initFileServerClient() error {
	fsClient, err := fsclient.WithCopyFsClient(*hc.FsClient)

	if err != nil {
		return err
	}

	hc.Uploader = upload.NewFileUploader(fsClient)
	return nil
}

func (hc *HTTPController) response(response interface{}) Response {
	return Response{
		Response: response,
	}
}

func (hc *HTTPController) message(message string) MessageResponse {
	return MessageResponse{
		Status:  200,
		Message: message,
	}
}
