package actions

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/lib"
	"github.com/zikwall/blogchain/src/app/models/user"
	"github.com/zikwall/blogchain/src/app/models/user/forms"
	"github.com/zikwall/blogchain/src/app/utils"
)

func (a BlogchainActionProvider) Logout(ctx *fiber.Ctx) error {
	return ctx.JSON(a.message("Successfully logout"))
}

type AuthResponse struct {
	Token string          `json:"token"`
	User  user.PublicUser `json:"user"`
}

func (a BlogchainActionProvider) Login(ctx *fiber.Ctx) error {
	form := &forms.LoginForm{}

	if err := ctx.BodyParser(&form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	if err := form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	result, err := user.ContextConnection(ctx.Context(), a.Db).FindByCredentials(form.Username)

	if err != nil {
		return exceptions.Wrap("failed check user", err)
	}

	if !result.Exist() || !utils.BlogchainPasswordCorrectness(result.PasswordHash, form.Password) {
		return exceptions.Wrap("login", errors.New("incorrect password was entered or the user doesn't exist."))
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 1000, a.RSA.GetPrivateKey())

	if err != nil {
		return exceptions.Wrap("invalid token", err)
	}

	return ctx.JSON(a.response(AuthResponse{
		Token: token,
		User:  result.Properties(),
	}))
}

func (a BlogchainActionProvider) Register(ctx *fiber.Ctx) error {
	form := &forms.RegisterForm{}

	if err := ctx.BodyParser(&form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	if err := form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	context := user.ContextConnection(ctx.Context(), a.Db)
	result, err := context.FindByUsernameOrEmail(form.Username, form.Email)

	if err != nil {
		return exceptions.Wrap("failed check user", err)
	}

	if result.Exist() {
		return exceptions.Wrap("register", errors.New("this name or email already exist."))
	}

	result, err = context.CreateUser(form)

	if err != nil {
		return exceptions.Wrap("failed create user", err)
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 100, a.RSA.GetPrivateKey())

	if err != nil {
		return exceptions.Wrap("invalid token", err)
	}

	return ctx.JSON(a.response(AuthResponse{
		Token: token,
		User:  result.Properties(),
	}))
}
