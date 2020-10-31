package actions

import (
	"github.com/gofiber/fiber/v2"
	content2 "github.com/zikwall/blogchain/src/models/content"
	"github.com/zikwall/blogchain/src/models/content/forms"
	"github.com/zikwall/blogchain/src/models/user"
	"strconv"
)

func GetEditContent(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	userInstance := c.Locals("user").(*user.User)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			//"status":  100,
			"message": "Content not found",
		})
	}

	content, err := content2.FindContentByIdAndUser(id, userInstance.Id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			//"status":  100,
			"message": "Content not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		//"status":  200,
		"content": content.ToJSONAPI(),
	})
}

func UpdateContent(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	userInstance := c.Locals("user").(*user.User)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})
	}

	form := &forms.ContentForm{
		UserId:  0,
		Title:   "",
		Content: "",
	}

	if err := c.BodyParser(form); err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Failed to parse your request body.",
		})
	}

	form.UserId = userInstance.Id

	if !form.Validate() {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Invalid request body fields.",
		})
	}

	content, err := content2.FindContentByIdAndUser(id, userInstance.Id)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})
	}

	img, err := c.FormFile("image")
	form.SetImage(forms.FormImage{img, err})

	err = content2.UpdateContent(content, form, c)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Что-то пошло не так...",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Успешно!",
	})
}

func AddContent(c *fiber.Ctx) error {
	userInstance := c.Locals("user").(*user.User)

	form := &forms.ContentForm{
		UserId:  0,
		Title:   "",
		Content: "",
	}

	if err := c.BodyParser(form); err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Failed to parse your request body.",
		})
	}

	form.UserId = userInstance.Id

	if !form.Validate() {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Invalid request body fields.",
		})
	}

	img, err := c.FormFile("image")
	form.SetImage(forms.FormImage{img, err})

	content, err := content2.CreateContent(form, c)
	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status":     200,
		"content_id": content.Id,
		"message":    "Successfully",
	})
}
