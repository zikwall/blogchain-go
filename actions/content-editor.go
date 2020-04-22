package actions

import (
	"fmt"
	"github.com/gofiber/fiber"
	content2 "github.com/zikwall/blogchain/models/content"
	"github.com/zikwall/blogchain/models/content/forms"
	"github.com/zikwall/blogchain/models/user"
	"strconv"
)

func GetEditContent(c *fiber.Ctx) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	userInstance := c.Locals("user").(*user.User)

	if err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})

		return
	}

	content, err := content2.FindContentByIdAndUser(id, userInstance.Id)

	if err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})

		return
	}

	c.JSON(fiber.Map{
		"status":  200,
		"content": content.ToJSONAPI(),
	})
}

func UpdateContent(c *fiber.Ctx) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	userInstance := c.Locals("user").(*user.User)

	if err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})

		return
	}

	form := &forms.ContentForm{
		UserId:  0,
		Title:   "",
		Content: "",
	}

	if err := c.BodyParser(form); err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Failed to parse your request body.",
		})

		return
	}

	form.UserId = userInstance.Id

	if !form.Validate() {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Invalid request body fields.",
		})

		return
	}

	content, err := content2.FindContentByIdAndUser(id, userInstance.Id)

	if err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})

		return
	}

	file, err := c.FormFile("image")
	if err == nil {
		form.Image = file.Filename
		_ = c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", file.Filename))
	}

	err = content2.UpdateContent(content, form)

	if err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Что-то пошло не так...",
		})

		return
	}

	c.JSON(fiber.Map{
		"status":  200,
		"message": "Успешно!",
	})
}

func AddContent(c *fiber.Ctx) {
	userInstance := c.Locals("user").(*user.User)

	form := &forms.ContentForm{
		UserId:  0,
		Title:   "",
		Content: "",
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

	form.UserId = userInstance.Id
	content, err := content2.CreateContent(form)
	if err != nil {
		panic(err)
	}

	c.JSON(fiber.Map{
		"status":     200,
		"content_id": content.Id,
		"message":    "Successfully",
	})
}
