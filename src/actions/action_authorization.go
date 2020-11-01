package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/user"
	"github.com/zikwall/blogchain/src/models/user/forms"
	"github.com/zikwall/blogchain/src/types"
	"github.com/zikwall/blogchain/src/utils"
)

func Logout(c *fiber.Ctx) error {
	return c.JSON(
		BlogchainMessageResponse{
			BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
				Status: 200,
			},
			Message: "Successfully logout",
		},
	)
}

func Login(c *fiber.Ctx) error {
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

	u := user.NewUserModel()
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

	token, err := types.CreateToken(result)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   result.Properties(),
	})
}

func Register(c *fiber.Ctx) error {
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

	u := user.NewUserModel()
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
		panic(err)
	}

	token, err := types.CreateToken(result)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   result.Properties(),
	})
}
