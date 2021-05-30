package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
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
		return exceptions.Wrap("failed parse content id", exceptions.NewErrApplicationLogic(err))
	}

	result, err := content.ContextConnection(ctx.Context(), a.Db).FindContentById(id)

	if err != nil {
		return exceptions.Wrap("failed find content by id", err)
	}

	viewers, err := statistic.GetPostViewersCount(a.StatsBatcher.Clickhouse, result.Id)

	if err != nil {
		log.Warning(err)
	}

	return ctx.Status(200).JSON(a.response(ContentResponse{
		Content: result.Response(),
		Viewers: viewers,
	}))
}

func (a BlogchainActionProvider) Contents(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")

	contents, err, count := content.ContextConnection(ctx.Context(), a.Db).FindAllContent(tag, getPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find contents", err)
	}

	return ctx.Status(200).JSON(a.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStats(a.StatsBatcher.Clickhouse, contents),
	}))
}

func (a BlogchainActionProvider) ContentsUser(ctx *fiber.Ctx) error {
	user, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("Failed parse user id", err)
	}

	contents, err, count := content.ContextConnection(ctx.Context(), a.Db).FindAllByUser(user, getPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find user contents by id", err)
	}

	return ctx.Status(200).JSON(a.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStats(a.StatsBatcher.Clickhouse, contents),
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
