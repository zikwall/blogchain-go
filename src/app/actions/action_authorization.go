package actions

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/lib"
	"github.com/zikwall/blogchain/src/app/models/user"
	"github.com/zikwall/blogchain/src/app/models/user/forms"
	"github.com/zikwall/blogchain/src/app/utils"
)

func (a BlogchainActionProvider) Logout(c *fiber.Ctx) error {
	return c.JSON(a.message("Successfully logout"))
}

func (a BlogchainActionProvider) Login(c *fiber.Ctx) error {
	form := &forms.LoginForm{}

	if err := c.BodyParser(&form); err != nil {
		return c.JSON(a.error(err))
	}

	if err := form.Validate(); err != nil {
		return c.JSON(err)
	}

	u := user.CreateUserConnection(a.db)
	result, err := u.FindByCredentials(form.Username)

	if err != nil {
		return c.JSON(a.error(err))
	}

	if !result.Exist() || !utils.BlogchainPasswordCorrectness(result.PasswordHash, form.Password) {
		return c.JSON(a.error(errors.New("An incorrect password was entered or the user does not exist.")))
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 1000, a.rsa.GetPrivateKey())

	if err != nil {
		return c.JSON(a.error(err))
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   result.Properties(),
	})
}

func (a BlogchainActionProvider) Register(c *fiber.Ctx) error {
	form := &forms.RegisterForm{}

	if err := c.BodyParser(&form); err != nil {
		return c.JSON(a.error(err))
	}

	if err := form.Validate(); err != nil {
		return c.JSON(a.error(err))
	}

	u := user.CreateUserConnection(a.db)
	result, err := u.FindByUsernameOrEmail(form.Username, form.Email)

	if err != nil {
		return c.JSON(a.error(err))
	}

	if result.Exist() {
		return c.JSON(a.error(errors.New("This name or email already exist.")))
	}

	result, err = u.CreateUser(form)

	if err != nil {
		return c.JSON(a.error(err))
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 100, a.rsa.GetPrivateKey())

	if err != nil {
		return c.JSON(a.error(err))
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   result.Properties(),
	})
}
