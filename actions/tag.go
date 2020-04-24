package actions

import (
	"github.com/gofiber/fiber"
	"github.com/zikwall/blogchain/models/tag"
)

func Tags(c *fiber.Ctx) {
	tags, _ := tag.GetTags()

	c.JSON(fiber.Map{
		"status": 200,
		"tags":   tags,
	})
}
