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

func (hc *HTTPController) Content(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse content id", exceptions.ThrowPublicError(err))
	}

	result, err := repositories.UseContentRepository(ctx.Context(), hc.DB).
		FindContentByID(id)

	if err != nil {
		return exceptions.Wrap("failed find content by id", err)
	}

	viewers, err := statistic.GetPostViewersCount(ctx.Context(), hc.Clickhouse, result.ID)

	if err != nil {
		log.Warning(err)
	}

	return ctx.Status(200).JSON(hc.response(ContentResponse{
		Content: result.Response(),
		Viewers: viewers,
	}))
}

func (hc *HTTPController) Contents(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")

	contents, count, err := repositories.UseContentRepository(ctx.Context(), hc.DB).
		FindAllContent(tag, extractPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find contents", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStatsContext(ctx.Context(), hc.Clickhouse, contents),
	}))
}

func (hc *HTTPController) ContentsUser(ctx *fiber.Ctx) error {
	user, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return exceptions.Wrap("failed parse user id", err)
	}

	contents, count, err := repositories.UseContentRepository(ctx.Context(), hc.DB).
		FindAllByUser(user, extractPageFromContext(ctx))

	if err != nil {
		return exceptions.Wrap("failed find user contents by id", err)
	}

	return ctx.Status(200).JSON(hc.response(ContentsResponse{
		Contents: contents,
		Meta: Meta{
			Pages: count,
		},
		Stats: withStatsContext(ctx.Context(), hc.Clickhouse, contents),
	}))
}

func withStatsContext(ctx context.Context, ch *clickhouse.Connection, cs []repositories.PublicContent) map[int64]uint64 {
	if len(cs) == 0 {
		return map[int64]uint64{}
	}

	ids := make([]int64, 0, len(cs))

	for i := range cs {
		ids = append(ids, cs[i].ID)
	}

	viewers, err := statistic.GetPostsViewersCount(ctx, ch, ids...)

	if err != nil {
		log.Warning(err)
	}

	return viewers
}
