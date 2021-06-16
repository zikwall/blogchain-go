package actions

import (
	"github.com/zikwall/blogchain/src/app/lib/upload"
	"github.com/zikwall/blogchain/src/app/statistic"
	"github.com/zikwall/blogchain/src/platform/container"
	"github.com/zikwall/blogchain/src/platform/database"
	"github.com/zikwall/blogchain/src/platform/maxmind"
)

// The HttpController structure is the base object for all http handlers,
// and encapsulates access to services such as databases, redis, etc.
type HttpController struct {
	RSA         container.RSA
	Db          *database.Connection
	StatsPacker *statistic.Packer
	Finder      *maxmind.Finder
	Uploader    upload.Uploader
}

func CreateHttpControllerWithCopy(p HttpController) *HttpController {
	return &HttpController{
		RSA:         p.RSA,
		Db:          p.Db,
		StatsPacker: p.StatsPacker,
		Finder:      p.Finder,
		Uploader:    p.Uploader,
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
