package actions

import (
	"github.com/gofiber/fiber/v2"
	"runtime"
	"time"
)

type Runtime struct {
	Metrics metrics `json:"metrics"`
	Uptime  uptime  `json:"uptime"`
}

type uptime struct {
	Seconds float64 `json:"timestamp"`
}

type metrics struct {
	MemoryAlloc uint64 `json:"memory_alloc"`
	Gorutines   int    `json:"gorutines"`
	NumGc       uint32 `json:"num_gc"`
}

func kb(b uint64) uint64 {
	return b / 1024
}

func RuntimeStatistic(startedAt time.Time) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		memory := runtime.MemStats{}
		runtime.ReadMemStats(&memory)

		return ctx.JSON(Runtime{
			Metrics: metrics{
				MemoryAlloc: kb(memory.Alloc),
				Gorutines:   runtime.NumGoroutine(),
				NumGc:       memory.NumGC,
			},
			Uptime: uptime{
				Seconds: time.Since(startedAt).Seconds(),
			},
		})
	}
}
