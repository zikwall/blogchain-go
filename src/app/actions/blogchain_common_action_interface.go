package actions

import (
	"github.com/gofiber/fiber/v2"
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
