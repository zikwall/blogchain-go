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

	model := content.CreateContentConnection(a.db)
	res, err := model.UserContent(id, userInstance.Id)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})
	}

	img, err := c.FormFile("image")
	form.SetImage(forms.FormImage{img, err})

	if err = model.UpdateContent(res, form, c); err != nil {
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

func (a BlogchainActionProvider) ContentCreate(c *fiber.Ctx) error {
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

	model := content.CreateContentConnection(a.db)
	result, err := model.CreateContent(form, c)

	if err != nil {
		return c.JSON(
			BlogchainMessageResponse{
				BlogchainCommonResponseAttributes: BlogchainCommonResponseAttributes{
					Status: 100,
				},
				Message: err.Error(),
			},
		)
	}

	return c.JSON(fiber.Map{
		"status":     200,
		"content_id": result.Id,
		"message":    "Successfully",
	})
}
