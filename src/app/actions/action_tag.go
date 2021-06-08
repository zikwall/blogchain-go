package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/models/tag"
)

type TagResponse struct {
	Tags []tag.Tag `json:"tags"`
}

func (a BlogchainActionProvider) Tags(ctx *fiber.Ctx) error {
	tags, err := tag.
		ContextConnection(ctx.Context(), a.Db).
		All()

	if err != nil {
		return exceptions.Wrap("failed find all tags", err)
	}

	return ctx.JSON(a.response(TagResponse{
		Tags: tags,
	}))
}
