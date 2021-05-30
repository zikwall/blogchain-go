package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/user"
	"strconv"
)

// ToDo: Need to bring all response formats to a single structure and format
type (
	MessageResponse struct {
		Status  uint8  `json:"status"`
		Message string `json:"message"`
	}
	Response struct {
		Response interface{} `json:"response"`
	}
)

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

func getUserFromContext(ctx *fiber.Ctx) *user.User {
	userInstance, ok := ctx.Locals("user").(*user.User)

	if !ok {
		return nil
	}

	return userInstance
}
