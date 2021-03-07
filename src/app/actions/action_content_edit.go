package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/content"
	"github.com/zikwall/blogchain/src/app/models/content/forms"
	"github.com/zikwall/blogchain/src/app/models/user"
	"strconv"
)

func (a BlogchainActionProvider) ContentInformation(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	userInstance := c.Locals("user").(*user.User)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			//"status":  100,
			"message": "Content not found",
		})
	}

	model := content.CreateContentConnection(a.db)
	result, err := model.UserContent(id, userInstance.Id)

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

func (a BlogchainActionProvider) ContentUpdate(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	userInstance := c.Locals("user").(*user.User)

	if err != nil {
		return c.JSON(a.error(err))
	}

	form := &forms.ContentForm{}

	if err := c.BodyParser(form); err != nil {
		return c.JSON(a.error(err))
	}

	form.UserId = userInstance.Id

	if err := form.Validate(); err != nil {
		return c.JSON(a.error(err))
	}

	model := content.CreateContentConnection(a.db)
	res, err := model.UserContent(id, userInstance.Id)

	if err != nil {
		return c.JSON(a.error(err))
	}

	img, err := c.FormFile("image")
	form.SetImage(forms.FormImage{img, err})

	if err := model.UpdateContent(res, form, c); err != nil {
		return c.JSON(a.error(err))
	}

	return c.JSON(a.message("Successfully!"))
}

func (a BlogchainActionProvider) ContentCreate(c *fiber.Ctx) error {
	userInstance := c.Locals("user").(*user.User)

	form := &forms.ContentForm{}

	if err := c.BodyParser(form); err != nil {
		return c.JSON(a.error(err))
	}

	form.UserId = userInstance.Id

	if err := form.Validate(); err != nil {
		return c.JSON(a.error(err))
	}

	img, err := c.FormFile("image")
	form.SetImage(forms.FormImage{img, err})

	model := content.CreateContentConnection(a.db)
	result, err := model.CreateContent(form, c)

	if err != nil {
		return c.JSON(a.error(err))
	}

	return c.JSON(fiber.Map{
		"status":     200,
		"content_id": result.Id,
		"message":    "Successfully",
	})
}
