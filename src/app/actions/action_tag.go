package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/tag"
)

type TagResponse struct {
	Tags []tag.Tag `json:"tags"`
}

func (a BlogchainActionProvider) Tags(ctx *fiber.Ctx) error {
	t := tag.CreateTagConnection(a.Db)
	tags, _ := t.All()

	return ctx.JSON(a.response(TagResponse{
		Tags: tags,
	}))
}
