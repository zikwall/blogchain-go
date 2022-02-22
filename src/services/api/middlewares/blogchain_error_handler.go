package middlewares

import (
	"errors"
	"fmt"

	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/pkg/log"
	"github.com/zikwall/blogchain/src/services/api/actions"

	"github.com/gofiber/fiber/v2"
)

const defaultErrorMessage = "Internal Server Error"

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		value := defaultErrorMessage

		var e *fiber.Error
		var w *exceptions.WrapError

		if errors.As(err, &e) {
			code = e.Code
			value = e.Message
		}

		if errors.As(err, &w) {
			var pub *exceptions.ErrPublic
			var pri *exceptions.ErrPrivate
			if errors.As(err, &pub) {
				value = fmt.Sprintf("%s: %v", w.Context, pub.Error())
			} else if errors.As(err, &pri) {
				log.Warning(fmt.Sprintf("%s: %v", w.Context, pri.Error()))
			}
		}

		return ctx.Status(code).JSON(actions.MessageResponse{
			Status:  100,
			Message: value,
		})
	}

	return nil
}
