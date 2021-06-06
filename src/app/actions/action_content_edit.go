package actions

import (
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/models/content"
	"github.com/zikwall/blogchain/src/app/models/content/forms"
	"github.com/zikwall/blogchain/src/app/utils"
	"strconv"
)

type (
	ContentCreatedResponse struct {
		ContentId int64 `json:"content_id"`
	}
)

func (a BlogchainActionProvider) ContentInformation(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse content id", exceptions.NewErrApplicationLogic(err))
	}

	result, err := content.ContextConnection(ctx.Context(), a.Db).UserContent(id, getUserFromContext(ctx).Id)

	if err != nil {
		return exceptions.Wrap("failed find user content", err)
	}

	return ctx.Status(200).JSON(a.response(ContentResponse{
		Content: result.Response(),
	}))
}

func (a BlogchainActionProvider) ContentUpdate(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse content id", exceptions.NewErrApplicationLogic(err))
	}

	form := &forms.ContentForm{}

	if err = ctx.BodyParser(form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	form.UserId = getUserFromContext(ctx).Id

	if err = form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	context := content.ContextConnection(ctx.Context(), a.Db)
	res, err := context.UserContent(id, form.UserId)

	if err != nil {
		return exceptions.Wrap("failed find user content", err)
	}

	if img, err := ctx.FormFile("image"); err == nil {
		filename := utils.CreateImagePath(res.Uuid)
		res.Image.String = filename

		file, err := img.Open()

		if err != nil {
			return exceptions.Wrap("failed open image file", err)
		}

		defer func() {
			_ = file.Close()
		}()

		if err := a.Uploader.UploadFile(ctx.Context(), filename, file); err != nil {
			return err
		}
	}

	if err := context.UpdateContent(res, form); err != nil {
		return exceptions.Wrap("failed update user content", err)
	}

	return ctx.Status(200).JSON(a.message("Successfully!"))
}

func (a BlogchainActionProvider) ContentCreate(ctx *fiber.Ctx) error {
	form := &forms.ContentForm{}

	if err := ctx.BodyParser(form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	form.UserId = getUserFromContext(ctx).Id
	form.UUID = uuid.NewV4().String()

	if err := form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	if img, err := ctx.FormFile("image"); err == nil {
		filename := utils.CreateImagePath(form.UUID)
		form.ImageName = filename

		file, err := img.Open()

		if err != nil {
			return exceptions.Wrap("failed open image file", err)
		}

		defer func() {
			_ = file.Close()
		}()

		if err := a.Uploader.UploadFile(ctx.Context(), filename, file); err != nil {
			return err
		}
	}

	result, err := content.ContextConnection(ctx.Context(), a.Db).CreateContent(form)

	if err != nil {
		return exceptions.Wrap("failed create user content", err)
	}

	return ctx.Status(200).JSON(a.response(ContentCreatedResponse{
		ContentId: result.Id,
	}))
}
