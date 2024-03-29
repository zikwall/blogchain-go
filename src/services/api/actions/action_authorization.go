package actions

import (
	"errors"

	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/pkg/jwt"
	"github.com/zikwall/blogchain/src/services/api/forms"
	"github.com/zikwall/blogchain/src/services/api/repositories"
	"github.com/zikwall/blogchain/src/services/api/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthResponse struct {
	Token string                  `json:"token"`
	User  repositories.PublicUser `json:"user"`
}

func (hc *HTTPController) Logout(ctx *fiber.Ctx) error {
	return ctx.JSON(hc.message("Successfully logout"))
}

func (hc *HTTPController) Login(ctx *fiber.Ctx) error {
	form := &forms.LoginForm{}

	if err := ctx.BodyParser(&form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	if err := form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	result, err := repositories.UseUserRepository(ctx.Context(), hc.DB).FindByCredentials(form.Username)
	if err != nil {
		return exceptions.Wrap("failed check user", err)
	}

	if !result.Exist() || !utils.CompareHashAndPassword(result.PasswordHash, form.Password) {
		return exceptions.Wrap("login", errors.New("incorrect password was entered or the user doesn't exist"))
	}

	claims := &jwt.TokenClaims{
		UUID: result.GetID(),
	}

	token, err := jwt.CreateJwtToken(claims, 1000, hc.RSA.GetPrivateKey())
	if err != nil {
		return exceptions.Wrap("invalid token", err)
	}

	return ctx.JSON(hc.response(AuthResponse{
		Token: token,
		User:  result.Properties(),
	}))
}

func (hc *HTTPController) Register(ctx *fiber.Ctx) error {
	form := &forms.RegisterForm{}

	if err := ctx.BodyParser(&form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	if err := form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	context := repositories.UseUserRepository(ctx.Context(), hc.DB)
	result, err := context.FindByUsernameOrEmail(form.Username, form.Email)
	if err != nil {
		return exceptions.Wrap("failed check user", err)
	}

	if result.Exist() {
		return exceptions.Wrap("register", errors.New("this name or email already exist"))
	}

	result, err = context.CreateUser(form)
	if err != nil {
		return exceptions.Wrap("failed create user", err)
	}

	claims := &jwt.TokenClaims{
		UUID: result.GetID(),
	}

	token, err := jwt.CreateJwtToken(claims, 100, hc.RSA.GetPrivateKey())
	if err != nil {
		return exceptions.Wrap("invalid token", err)
	}

	return ctx.JSON(hc.response(AuthResponse{
		Token: token,
		User:  result.Properties(),
	}))
}
