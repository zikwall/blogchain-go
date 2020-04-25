package actions

import (
	"github.com/gofiber/fiber"
	content2 "github.com/zikwall/blogchain/models/content"
	"strconv"
)

func GetContent(c *fiber.Ctx) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})

		return
	}

	content, err := content2.FindContentById(id)

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

func GetContents(c *fiber.Ctx) {
	tag := c.Params("tag")

	contents, err := content2.FindAllContent(tag)
	if err != nil {
		c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})

		return
	}

	c.JSON(fiber.Map{
		"status":   200,
		"contents": contents,
	})
}
