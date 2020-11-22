package utils

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
)

func SaveFile(ctx *fiber.Ctx, file *multipart.FileHeader, path string) error {
	return ctx.SaveFile(file, path)
}
