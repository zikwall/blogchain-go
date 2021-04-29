package actions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/statistic"
	"github.com/zikwall/blogchain/src/app/utils"
	"time"
)

func (a *BlogchainActionProvider) PushPostStats(ctx *fiber.Ctx) error {
	data := &statistic.PostStats{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.JSON(a.error(err))
	}

	now := time.Now()
	ip := fmt.Sprintf("%v", ctx.Locals("ip"))

	stats := statistic.PostStats{
		PostId:   data.PostId,
		OwnerId:  data.PostId,
		Os:       data.Os,
		Browser:  data.Browser,
		Platform: data.Platform,
		Ip:       ip,
		Country:  "",
		Region:   "",
		InsertTs: utils.Datetime(now),
		Date:     utils.Date(now),
	}

	geo, err := a.finder.Lookup(ip)

	if err == nil {
		stats.Country = geo.Country
		stats.Region = geo.Region
	}

	a.statsBatcher.AppendRecords(stats)

	return ctx.Status(200).SendString("OK")
}
