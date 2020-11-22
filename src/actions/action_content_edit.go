package actions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/models/content"
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

	model := content.NewContentModel()
	result, err := model.UserContent(id, userInstance.Id)

	fmt.Println(userInstance.Id)
	fmt.Println(id)

	fmt.Println(result)

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

	model := content.NewContentModel()
	res, err := model.UserContent(id, userInstance.Id)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  100,
			"message": "Content not found",
		})
	}

	img, err := c.FormFile("image")
	form.SetImage(forms.FormImage{img, err})

	err = model.UpdateContent(res.Content, form, c)

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

	model := content.NewContentModel()
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
