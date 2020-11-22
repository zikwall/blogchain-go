package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/tag"
)

func (a BlogchainActionProvider) Tags(c *fiber.Ctx) error {
	t := tag.NewTagModel(a.db)
	tags, _ := t.All()

	return c.JSON(fiber.Map{
		"status": 200,
		"tags":   tags,
	})
}
