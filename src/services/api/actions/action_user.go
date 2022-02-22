package actions

import (
	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/services/api/repositories"

	"github.com/gofiber/fiber/v2"
)

type ProfileResponse struct {
	User repositories.PublicUser `json:"user"`
}

func (hc *HTTPController) Profile(ctx *fiber.Ctx) error {
	result, err := repositories.UseUserRepository(ctx.Context(), hc.DB).FindByUsername(ctx.Params("username"))
	if err != nil {
		return exceptions.Wrap("failed find user", err)
	}

	return ctx.Status(200).JSON(hc.response(ProfileResponse{
		User: result.Properties(),
	}))
}
