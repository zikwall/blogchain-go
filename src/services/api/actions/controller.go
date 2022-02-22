package actions

import (
	"github.com/zikwall/blogchain/src/pkg/clickhouse"
	"github.com/zikwall/blogchain/src/pkg/container"
	"github.com/zikwall/blogchain/src/pkg/database"
	"github.com/zikwall/blogchain/src/pkg/maxmind"
	"github.com/zikwall/blogchain/src/protobuf/storage"
	"github.com/zikwall/blogchain/src/services/api/lib/upload"

	"github.com/zikwall/fsclient"
)

// The HTTPController structure is the base object for all http handlers,
// and encapsulates access to services such as databases, redis, etc.
type HTTPController struct {
	RSA              container.RSA
	DB               *database.Connection
	Clickhouse       *clickhouse.Connection
	ClickhouseBuffer *clickhouse.BufferAdapter
	GeoReader        *maxmind.Reader
	Uploader         upload.Uploader
	FsClient         *fsclient.FsClient
	StatisticClient  storage.StorageClient
}

func CreateHTTPControllerWithCopy(p *HTTPController) (*HTTPController, error) {
	controller := &HTTPController{
		RSA:              p.RSA,
		DB:               p.DB,
		Clickhouse:       p.Clickhouse,
		ClickhouseBuffer: p.ClickhouseBuffer,
		GeoReader:        p.GeoReader,
		FsClient:         p.FsClient,
		StatisticClient:  p.StatisticClient,
	}

	if err := controller.after(); err != nil {
		return nil, err
	}

	return controller, nil
}

func (hc *HTTPController) after() error {
	return hc.initFileServerClient()
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
