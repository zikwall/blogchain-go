package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/models/content"
	"github.com/zikwall/blogchain/src/app/statistic"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/log"
	"strconv"
)

type (
	ContentResponse struct {
		Content content.PublicContent `json:"content"`
		Viewers uint64                `json:"viewers"`
	}
	ContentsResponse struct {
		Contents []content.PublicContent `json:"contents"`
		Meta     Meta                    `json:"meta"`
		Stats    map[int64]uint64        `json:"stats"`
	}
	Meta struct {
		Pages float64 `json:"pages"`
	}
)

func (a BlogchainActionProvider) Content(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(500).JSON(a.error(err))
	}

	model := content.CreateContentConnection(a.db)
	result, err := model.FindContentById(id)

	if err != nil {
		return ctx.Status(404).JSON(a.error(err))
	}

	viewers, err := statistic.GetPostViewersCount(a.statsBatcher.Clickhouse, result.Id)

	if err != nil {
		log.Warning(err)
	}

	// current view
	viewers += 1

	return ctx.Status(200).JSON(a.response(ContentResponse{
		Content: result.Response(),
		Viewers: viewers,
	}))
}

func (a BlogchainActionProvider) Contents(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")
	var page int64

	if ctx.Params("page") != "" {
		if p, err := strconv.ParseInt(ctx.Params("page"), 10, 64); err == nil {
			// client page 1 === 0 in server side
			page = p - 1
		}
	}

	model := content.CreateContentConnection(a.db)
	contents, err, count := model.FindAllContent(tag, page)

	if err != nil {
		return ctx.Status(404).JSON(a.error(err))
	}

	return ctx.Status(200).JSON(a.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStats(a.statsBatcher.Clickhouse, contents),
	}))
}

func (a BlogchainActionProvider) ContentsUser(ctx *fiber.Ctx) error {
	user, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(404).JSON(a.error(err))
	}

	var page int64

	if ctx.Params("page") != "" {
		if p, err := strconv.ParseInt(ctx.Params("page"), 10, 64); err == nil {
			// client page 1 === 0 in server side
			page = p - 1
		}
	}

	model := content.CreateContentConnection(a.db)
	contents, err, count := model.FindAllByUser(user, page)

	if err != nil {
		return ctx.Status(404).JSON(a.error(err))
	}

	return ctx.Status(200).JSON(a.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStats(a.statsBatcher.Clickhouse, contents),
	}))
}

func withStats(ch *clickhouse.Clickhouse, contents []content.PublicContent) map[int64]uint64 {
	var ids []int64

	for _, c := range contents {
		ids = append(ids, c.Id)
	}

	viewers, err := statistic.GetPostsViewersCount(ch, ids...)

	if err != nil {
		log.Warning(err)
	}

	return viewers
}
