package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/content"
	"strconv"
)

func (a BlogchainActionProvider) Content(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			//"status":  100,
			"message": "Content not found",
		})
	}

	model := content.NewContentModel(a.db)
	result, err := model.FindContentById(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			//"status":  100,
			"message": "Content not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		//"status":  200,
		"content": result.Response(),
	})
}

func (a BlogchainActionProvider) Contents(c *fiber.Ctx) error {
	tag := c.Params("tag")
	var page int64

	if c.Params("page") != "" {
		if p, err := strconv.ParseInt(c.Params("page"), 10, 64); err == nil {
			page = p
		}
	}

	model := content.NewContentModel(a.db)
	contents, err, count := model.FindAllContent(tag, page)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			//"status":  100,
			"message": "Content not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		//"status":   200,
		"contents": contents,
		"meta": fiber.Map{
			"pages": count,
		},
	})
}

func (a BlogchainActionProvider) ContentsUser(c *fiber.Ctx) error {
	user, err := strconv.ParseInt(c.Params("id"), 10, 64)
	var page int64

	if c.Params("page") != "" {
		if p, err := strconv.ParseInt(c.Params("page"), 10, 64); err == nil {
			page = p
		}
	}

	model := content.NewContentModel(a.db)
	contents, err, count := model.FindAllByUser(user, page)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Content not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"contents": contents,
		"meta": fiber.Map{
			"pages": count,
		},
	})
}
