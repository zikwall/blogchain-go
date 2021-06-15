package actions

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/repositories"
	"github.com/zikwall/blogchain/src/app/statistic"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/log"
	"strconv"
)

type ContentResponse struct {
	Content repositories.PublicContent `json:"content"`
	Viewers uint64                     `json:"viewers"`
}

type ContentsResponse struct {
	Contents []repositories.PublicContent `json:"contents"`
	Meta     Meta                         `json:"meta"`
	Stats    map[int64]uint64             `json:"stats"`
}

type Meta struct {
	Pages float64 `json:"pages"`
}

func (hc HttpController) Content(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse content id", exceptions.NewErrApplicationLogic(err))
	}

	result, err := repositories.UseContentRepository(ctx.Context(), hc.Db).
		FindContentById(id)

	if err != nil {
		return exceptions.Wrap("failed find content by id", err)
	}

	viewers, err := statistic.GetPostViewersCount(ctx.Context(), hc.StatsPacker.Clickhouse, result.Id)

	if err != nil {
		log.Warning(err)
	}

	return ctx.Status(200).JSON(hc.response(ContentResponse{
		Content: result.Response(),
		Viewers: viewers,
	}))
}

func (hc HttpController) Contents(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")

	contents, err, count := repositories.UseContentRepository(ctx.Context(), hc.Db).
		FindAllContent(tag, extractPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find contents", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStatsContext(ctx.Context(), hc.StatsPacker.Clickhouse, contents),
	}))
}

func (hc HttpController) ContentsUser(ctx *fiber.Ctx) error {
	user, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse user id", err)
	}

	contents, err, count := repositories.UseContentRepository(ctx.Context(), hc.Db).
		FindAllByUser(user, extractPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find user contents by id", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStatsContext(ctx.Context(), hc.StatsPacker.Clickhouse, contents),
	}))
}

func withStatsContext(context context.Context, ch *clickhouse.Clickhouse, cs []repositories.PublicContent) map[int64]uint64 {
	if len(cs) == 0 {
		return map[int64]uint64{}
	}

	ids := make([]int64, 0, len(cs))

	for _, c := range cs {
		ids = append(ids, c.Id)
	}

	viewers, err := statistic.GetPostsViewersCount(context, ch, ids...)

	if err != nil {
		log.Warning(err)
	}

	return viewers
}
