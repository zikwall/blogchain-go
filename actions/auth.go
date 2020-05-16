package actions

import (
	"github.com/gofiber/fiber"
	user2 "github.com/zikwall/blogchain/models/user"
	"github.com/zikwall/blogchain/models/user/forms"
	"github.com/zikwall/blogchain/types"
)

func Logout(c *fiber.Ctx) {
	c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Successful",
	})
}

func Login(c *fiber.Ctx) {
	form := &forms.LoginForm{
		Username: "",
		Password: "",
	}

	if err := c.BodyParser(&form); err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Failed to parse your request body.",
		})

		return
	}

	if !form.Validate() {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Invalid request body fields.",
		})

		return
	}

	user, _ := user2.FindByCredentials(form.Username)

	if !user.Exist() || !user2.PasswordFirewall(user.PasswordHash, form.Password) {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Wrong password or user not found.",
		})

		return
	}

	token, err := types.CreateToken(user)

	if err != nil {
		panic(err)
	}

	c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   user.Properties(),
	})
}

func Register(c *fiber.Ctx) {
	form := &forms.RegisterForm{
		Email:          "",
		Username:       "",
		Password:       "",
		PasswordRepeat: "",
		Name:           "",
	}

	if err := c.BodyParser(&form); err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Failed to parse your request body.",
		})

		return
	}

	if !form.Validate() || !form.ComparePasswords() {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Invalid request body fields.",
		})

		return
	}

	user, _ := user2.FindByUsernameOrEmail(form.Username, form.Email)

	if user.Exist() {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "This name or email already exist.",
		})

		return
	}

	user, err := user2.CreateUser(form)

	if err != nil {
		panic(err)
	}

	token, err := types.CreateToken(user)

	if err != nil {
		panic(err)
	}

	c.JSON(fiber.Map{
		"status": 200,
		"token":  token,
		"user":   user.Properties(),
	})
}
