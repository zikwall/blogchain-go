package actions

import (
	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/services/api/repositories"

	"github.com/gofiber/fiber/v2"
)

type TagResponse struct {
	Tags []repositories.Tag `json:"tags"`
}

func (hc *HTTPController) Tags(ctx *fiber.Ctx) error {
	tags, err := repositories.UseTagRepository(ctx.Context(), hc.DB).All()
	if err != nil {
		return exceptions.Wrap("failed find all tags", err)
	}

	return ctx.JSON(hc.response(TagResponse{
		Tags: tags,
	}))
}
