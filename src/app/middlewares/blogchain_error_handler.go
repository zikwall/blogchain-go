package middlewares

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/actions"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/platform/log"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		value := "Internal Server Error"

		var e *fiber.Error
		var wrap *exceptions.WrapError

		if errors.As(err, &e) {
			code = e.Code
			value = e.Message
		}

		if errors.As(err, &wrap) {
			var appErr *exceptions.ErrApplicationLogic
			var dbErr *exceptions.ErrDatabaseAccess

			if errors.As(err, &appErr) {
				value = fmt.Sprintf("%s: %v", wrap.Context, appErr.Error())
			} else if errors.As(err, &dbErr) {
				log.Warning(fmt.Sprintf("%s: %v", wrap.Context, dbErr.Error()))
			}
		}

		return ctx.Status(code).JSON(actions.MessageResponse{
			Status:  100,
			Message: value,
		})
	}

	return nil
}
