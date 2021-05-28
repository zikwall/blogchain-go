package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/content"
	"github.com/zikwall/blogchain/src/app/models/content/forms"
	"github.com/zikwall/blogchain/src/app/models/user"
	"strconv"
)

type (
	ContentCreatedResponse struct {
		ContentId int64 `json:"content_id"`
	}
)

func (a BlogchainActionProvider) ContentInformation(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	userInstance := ctx.Locals("user").(*user.User)

	if err != nil {
		return ctx.Status(500).JSON(a.response(err))
	}

	model := content.CreateContentConnection(a.db)
	result, err := model.UserContent(id, userInstance.Id)

	if err != nil {
		return ctx.Status(404).JSON(a.response(err))
	}

	return ctx.Status(200).JSON(a.response(ContentResponse{
		Content: result.Response(),
	}))
}

func (a BlogchainActionProvider) ContentUpdate(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	userInstance := ctx.Locals("user").(*user.User)

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	form := &forms.ContentForm{}

	if err = ctx.BodyParser(form); err != nil {
		return ctx.JSON(a.error(err))
	}

	form.UserId = userInstance.Id

	if err = form.Validate(); err != nil {
		return ctx.JSON(a.error(err))
	}

	model := content.CreateContentConnection(a.db)
	res, err := model.UserContent(id, userInstance.Id)

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	img, err := ctx.FormFile("image")
	form.SetImage(forms.FormImage{File: img, Err: err})

	if err := model.UpdateContent(res, form, ctx); err != nil {
		return ctx.JSON(a.error(err))
	}

	return ctx.Status(200).JSON(a.message("Successfully!"))
}

func (a BlogchainActionProvider) ContentCreate(ctx *fiber.Ctx) error {
	userInstance := ctx.Locals("user").(*user.User)

	form := &forms.ContentForm{}

	if err := ctx.BodyParser(form); err != nil {
		return ctx.JSON(a.error(err))
	}

	form.UserId = userInstance.Id

	if err := form.Validate(); err != nil {
		return ctx.JSON(a.error(err))
	}

	img, err := ctx.FormFile("image")
	form.SetImage(forms.FormImage{File: img, Err: err})

	model := content.CreateContentConnection(a.db)
	result, err := model.CreateContent(form, ctx)

	if err != nil {
		return ctx.JSON(a.error(err))
	}

	return ctx.Status(200).JSON(a.response(ContentCreatedResponse{
		ContentId: result.Id,
	}))
}
