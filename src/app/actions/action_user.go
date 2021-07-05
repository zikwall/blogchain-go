package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/repositories"
)

type ProfileResponse struct {
	User repositories.PublicUser `json:"user"`
}

func (hc *HTTPController) Profile(ctx *fiber.Ctx) error {
	result, err := repositories.UseUserRepository(ctx.Context(), hc.Db).FindByUsername(ctx.Params("username"))

	if err != nil {
		return exceptions.Wrap("failed find user", err)
	}

	return ctx.Status(200).JSON(hc.response(ProfileResponse{
		User: result.Properties(),
	}))
}
