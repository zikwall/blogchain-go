package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/content"
	"strconv"
)

type (
	ContentResponse struct {
		Content content.PublicContent `json:"content"`
	}
	ContentsResponse struct {
		Contents []content.PublicContent `json:"contents"`
		Meta     Meta                    `json:"meta"`
	}
	Meta struct {
		Pages float64 `json:"pages"`
	}
)

func (a BlogchainActionProvider) Content(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(500).JSON(a.error(err))
	}

	model := content.CreateContentConnection(a.db)
	result, err := model.FindContentById(id)

	if err != nil {
		return c.Status(404).JSON(a.error(err))
	}

	return c.Status(200).JSON(a.response(ContentResponse{
		Content: result.Response(),
	}))
}

func (a BlogchainActionProvider) Contents(c *fiber.Ctx) error {
	tag := c.Params("tag")
	var page int64

	if c.Params("page") != "" {
		if p, err := strconv.ParseInt(c.Params("page"), 10, 64); err == nil {
			// client page 1 === 0 in server side
			page = p - 1
		}
	}

	model := content.CreateContentConnection(a.db)
	contents, err, count := model.FindAllContent(tag, page)

	if err != nil {
		return c.Status(404).JSON(a.error(err))
	}

	return c.Status(200).JSON(a.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
	}))
}

func (a BlogchainActionProvider) ContentsUser(c *fiber.Ctx) error {
	user, err := strconv.ParseInt(c.Params("id"), 10, 64)
	var page int64

	if c.Params("page") != "" {
		if p, err := strconv.ParseInt(c.Params("page"), 10, 64); err == nil {
			// client page 1 === 0 in server side
			page = p - 1
		}
	}

	model := content.CreateContentConnection(a.db)
	contents, err, count := model.FindAllByUser(user, page)

	if err != nil {
		return c.Status(404).JSON(a.error(err))
	}

	return c.Status(200).JSON(a.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
	}))
}
