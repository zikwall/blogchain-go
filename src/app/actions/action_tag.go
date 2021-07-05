package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/repositories"
)

type TagResponse struct {
	Tags []repositories.Tag `json:"tags"`
}

func (hc *HTTPController) Tags(ctx *fiber.Ctx) error {
	tags, err := repositories.UseTagRepository(ctx.Context(), hc.Db).All()

	if err != nil {
		return exceptions.Wrap("failed find all tags", err)
	}

	return ctx.JSON(hc.response(TagResponse{
		Tags: tags,
	}))
}
