package actions

import (
	"github.com/gofiber/fiber/v2"
	user2 "github.com/zikwall/blogchain/src/models/user"
	"github.com/zikwall/blogchain/src/models/user/forms"
	"github.com/zikwall/blogchain/src/types"
)

func Logout(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Successful",
	})
}

func Login(c *fiber.Ctx) error {
	form := &forms.LoginForm{
		Username: "",
		Password: "",
	}

	if err := c.BodyParser(&form); err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Failed to parse your request body.",
		})
	}

	if !form.Validate() {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Invalid request body fields.",
		})
	}

	user, _ := user2.FindByCredentials(form.Username)

	if !user.Exist() || !user2.PasswordFirewall(user.PasswordHash, form.Password) {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Wrong password or user not found.",
		})
	}

	token, err := types.CreateToken(user)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   user.Properties(),
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
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Failed to parse your request body.",
		})
	}

	if !form.Validate() || !form.ComparePasswords() {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Invalid request body fields.",
		})
	}

	user, _ := user2.FindByUsernameOrEmail(form.Username, form.Email)

	if user.Exist() {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "This name or email already exist.",
		})
	}

	user, err := user2.CreateUser(form)

	if err != nil {
		panic(err)
	}

	token, err := types.CreateToken(user)

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   user.Properties(),
	})
}
