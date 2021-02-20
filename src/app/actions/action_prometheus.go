package actions

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusWithFastHTTPAdapter() fiber.Handler {
	return adaptor.HTTPHandler(promhttp.Handler())
}
