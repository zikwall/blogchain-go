package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/models/user"
)

type ProfileResponse struct {
	User user.PublicUser `json:"user"`
}

func (a BlogchainActionProvider) Profile(ctx *fiber.Ctx) error {
	u := user.CreateUserConnection(ctx.Context(), a.Db)

	result, err := u.FindByUsername(ctx.Params("username"))

	if err != nil {
		return exceptions.Wrap("failed find user", err)
	}

	return ctx.Status(200).JSON(a.response(ProfileResponse{
		User: result.Properties(),
	}))
}
