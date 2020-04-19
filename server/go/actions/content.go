package actions

import (
	"github.com/gofiber/fiber"
	content2 "github.com/zikwall/blogchain/models/content"
	"github.com/zikwall/blogchain/models/content/forms"
	"strconv"
)

func AddContent(c *fiber.Ctx) {
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

func GetContent(c *fiber.Ctx) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err == nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})
	}

	content, err := content2.FindContentById(id)

	if err != nil {
		//panic(err)
	}

	c.JSON(fiber.Map{
		"status":  200,
		"title":   content.Title,
		"content": content.Content,
		"user":    content.User.Properties(),
	})
}
