package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/repositories"
	"strconv"
)

type MessageResponse struct {
	Status  uint8  `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Response interface{} `json:"response"`
}

func getPageFromContext(ctx *fiber.Ctx) int64 {
	var page int64

	if ctx.Params("page") != "" {
		if p, err := strconv.ParseInt(ctx.Params("page"), 10, 64); err == nil {
			// client page 1 === 0 in server side
			page = p - 1
		}
	}

	return page
}

func getUserFromContext(ctx *fiber.Ctx) *repositories.User {
	userInstance, ok := ctx.Locals("user").(*repositories.User)

	if !ok {
		return nil
	}

	return userInstance
}
