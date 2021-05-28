package actions

import (
	"errors"
	"github.com/gofiber/fiber/v2"
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
		return ctx.JSON(a.error(err))
	}

	if err := form.Validate(); err != nil {
		return ctx.JSON(err)
	}

	u := user.CreateUserConnection(a.Db)
	result, err := u.FindByCredentials(form.Username)

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	if !result.Exist() || !utils.BlogchainPasswordCorrectness(result.PasswordHash, form.Password) {
		return ctx.JSON(a.error(errors.New("An incorrect password was entered or the user does not exist.")))
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 1000, a.RSA.GetPrivateKey())

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	return ctx.JSON(a.response(AuthResponse{
		Token: token,
		User:  result.Properties(),
	}))
}

func (a BlogchainActionProvider) Register(ctx *fiber.Ctx) error {
	form := &forms.RegisterForm{}

	if err := ctx.BodyParser(&form); err != nil {
		return ctx.JSON(a.error(err))
	}

	if err := form.Validate(); err != nil {
		return ctx.JSON(a.error(err))
	}

	u := user.CreateUserConnection(a.Db)
	result, err := u.FindByUsernameOrEmail(form.Username, form.Email)

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	if result.Exist() {
		return ctx.JSON(a.error(errors.New("This name or email already exist.")))
	}

	result, err = u.CreateUser(form)

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 100, a.RSA.GetPrivateKey())

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	return ctx.JSON(a.response(AuthResponse{
		Token: token,
		User:  result.Properties(),
	}))
}
