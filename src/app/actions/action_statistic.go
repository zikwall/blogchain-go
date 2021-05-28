package actions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mssola/user_agent"
	"github.com/zikwall/blogchain/src/app/statistic"
	"github.com/zikwall/blogchain/src/app/utils"
	"github.com/zikwall/blogchain/src/platform/maxmind"
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
		OwnerId:  data.OwnerId,
		Os:       "",
		Browser:  "",
		Platform: "",
		Ip:       ip,
		Country:  "",
		Region:   "",
		InsertTs: utils.Datetime(now),
		Date:     utils.Date(now),
	}

	geo, err := a.Finder.Lookup(ip)

	if err == nil {
		stats = withFinderAttributes(stats, geo)
	}

	if userAgent := ctx.Get("User-Agent", ""); userAgent != "" {
		stats = withUserAgent(stats, userAgent)
	}

	a.StatsBatcher.AppendRecords(stats)

	return ctx.Status(200).SendString("OK")
}

func withFinderAttributes(stats statistic.PostStats, result maxmind.FindResult) statistic.PostStats {
	stats.Region = result.Region
	stats.Country = result.Country

	return stats
}

func withUserAgent(stats statistic.PostStats, userAgent string) statistic.PostStats {
	ua := user_agent.New(userAgent)

	stats.Os = ua.OS()
	browser, version := ua.Browser()

	if browser != "" {
		stats.Browser = fmt.Sprintf("%s/%s", browser, version)
	}

	stats.Platform = ua.Platform()

	return stats
}
