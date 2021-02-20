package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/lib"
	"github.com/zikwall/blogchain/src/app/models/user"
	"github.com/zikwall/blogchain/src/app/models/user/forms"
	"github.com/zikwall/blogchain/src/app/utils"
)

func (a BlogchainActionProvider) Logout(c *fiber.Ctx) error {
	return c.JSON(
		BlogchainMessageResponse{
			BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
				Status: 200,
			},
			Message: "Successfully logout",
		},
	)
}

func (a BlogchainActionProvider) Login(c *fiber.Ctx) error {
	form := &forms.LoginForm{
		Username: "",
		Password: "",
	}

	if err := c.BodyParser(&form); err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: "Request could not be processed.",
			},
		)
	}

	if !form.Validate() {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: "Invalid data entered.",
			},
		)
	}

	u := user.NewUserModel(a.db)
	result, err := u.FindByCredentials(form.Username)

	if err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: err.Error(),
			},
		)
	}

	if !result.Exist() || !utils.BlogchainPasswordCorrectness(result.PasswordHash, form.Password) {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: "An incorrect password was entered or the user does not exist.",
			},
		)
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 1000, a.rsa.GetPrivateKey())

	if err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: err.Error(),
			},
		)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   result.Properties(),
	})
}

func (a BlogchainActionProvider) Register(c *fiber.Ctx) error {
	form := &forms.RegisterForm{
		Email:          "",
		Username:       "",
		Password:       "",
		PasswordRepeat: "",
		Name:           "",
	}

	if err := c.BodyParser(&form); err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: "Request could not be processed.",
			},
		)
	}

	if !form.Validate() || !form.ComparePasswords() {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: "Invalid data entered.",
			},
		)
	}

	u := user.NewUserModel(a.db)
	result, err := u.FindByUsernameOrEmail(form.Username, form.Email)

	if err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: err.Error(),
			},
		)
	}

	if result.Exist() {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: "This name or email already exist.",
			},
		)
	}

	result, err = u.CreateUser(form)

	if err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: "Internal server error.",
			},
		)
	}

	claims := lib.TokenClaims{
		UUID: result.GetId(),
	}

	token, err := lib.CreateJwtToken(claims, 100, a.rsa.GetPrivateKey())

	if err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: err.Error(),
			},
		)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   result.Properties(),
	})
}
