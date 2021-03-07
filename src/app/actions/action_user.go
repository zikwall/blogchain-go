package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/user"
)

type ProfileResponse struct {
	User user.PublicUser `json:"user"`
}

func (a BlogchainActionProvider) Profile(ctx *fiber.Ctx) error {
	u := user.CreateUserConnection(a.db)

	result, err := u.FindByUsername(ctx.Params("username"))

	if err != nil {
		return ctx.Status(404).JSON(a.error(err))
	}

	return ctx.Status(200).JSON(a.response(ProfileResponse{
		User: result.Properties(),
	}))
}
