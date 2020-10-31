package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/tag"
)

func Tags(c *fiber.Ctx) error {
	tags, _ := tag.GetTags()

	return c.JSON(fiber.Map{
		"status": 200,
		"tags":   tags,
	})
}
