package actions

import (
	"fmt"
	"time"

	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/pkg/maxmind"
	"github.com/zikwall/blogchain/src/services/api/statistic"
	"github.com/zikwall/blogchain/src/services/api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/mssola/user_agent"
)

func (hc *HTTPController) PushPostStats(ctx *fiber.Ctx) error {
	data := &statistic.PostStats{}
	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	now := time.Now()
	ip := fmt.Sprintf("%v", ctx.Locals("ip"))

	stats := &statistic.PostStats{
		PostID:   data.PostID,
		OwnerID:  data.OwnerID,
		Os:       "",
		Browser:  "",
		Platform: "",
		IP:       ip,
		Country:  "",
		Region:   "",
		InsertTS: utils.Datetime(now),
		Date:     utils.Date(now),
	}
	geo, err := hc.GeoReader.Lookup(ip)
	if err == nil {
		stats = withFinderAttributes(stats, geo)
	}
	if userAgent := ctx.Get("User-Agent", ""); userAgent != "" {
		stats = withUserAgent(stats, userAgent)
	}
	hc.writeAPI.WriteRow(stats)
	return ctx.Status(200).SendString("OK")
}

func withFinderAttributes(stats *statistic.PostStats, result maxmind.ReaderResult) *statistic.PostStats {
	stats.Region = result.Region
	stats.Country = result.Country
	return stats
}

func withUserAgent(stats *statistic.PostStats, userAgent string) *statistic.PostStats {
	ua := user_agent.New(userAgent)
	stats.Os = ua.OS()
	browser, version := ua.Browser()
	if browser != "" {
		stats.Browser = fmt.Sprintf("%s/%s", browser, version)
	}
	stats.Platform = ua.Platform()
	return stats
}
