package actions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/forms"
	"github.com/zikwall/blogchain/src/app/repositories"
	"strconv"
)

type ContentCreatedResponse struct {
	ContentId int64 `json:"content_id"`
}

func (hc *HTTPController) ContentInformation(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse content id", exceptions.ThrowPublicError(err))
	}

	result, err := repositories.UseContentRepository(ctx.Context(), hc.Db).
		UserContent(id, extractUserFromContext(ctx).Id)

	if err != nil {
		return exceptions.Wrap("failed find user content", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentResponse{
		Content: result.Response(),
	}))
}

func (hc *HTTPController) ContentUpdate(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse content id", exceptions.ThrowPublicError(err))
	}

	form := &forms.ContentForm{}

	if err = ctx.BodyParser(form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	form.UserId = extractUserFromContext(ctx).Id

	if err = form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	context := repositories.UseContentRepository(ctx.Context(), hc.Db)
	res, err := context.UserContent(id, form.UserId)

	if err != nil {
		return exceptions.Wrap("failed find user content", err)
	}

	if img, err := ctx.FormFile("image"); err == nil {
		filename := createImagePath(res.Uuid)
		res.Image.String = filename

		file, err := img.Open()

		if err != nil {
			return exceptions.Wrap("failed open image file", err)
		}

		defer func() {
			_ = file.Close()
		}()

		if err := hc.Uploader.UploadFile(ctx.Context(), filename, file); err != nil {
			return err
		}
	}

	if err := context.UpdateContent(&res, form); err != nil {
		return exceptions.Wrap("failed update user content", err)
	}

	return ctx.Status(200).JSON(hc.message("Successfully!"))
}

func (hc *HTTPController) ContentCreate(ctx *fiber.Ctx) error {
	form := &forms.ContentForm{}

	if err := ctx.BodyParser(form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	form.UserId = extractUserFromContext(ctx).Id
	form.UUID = uuid.NewV4().String()

	if err := form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	if img, err := ctx.FormFile("image"); err == nil {
		filename := createImagePath(form.UUID)
		form.ImageName = filename

		file, err := img.Open()

		if err != nil {
			return exceptions.Wrap("failed open image file", err)
		}

		defer func() {
			_ = file.Close()
		}()

		if err := hc.Uploader.UploadFile(ctx.Context(), filename, file); err != nil {
			return err
		}
	}

	result, err := repositories.UseContentRepository(ctx.Context(), hc.Db).CreateContent(form)

	if err != nil {
		return exceptions.Wrap("failed create user content", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentCreatedResponse{
		ContentId: result.Id,
	}))
}

func createImagePath(uuidv4 string) string {
	return fmt.Sprintf("%s.png", uuidv4)
}
