package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/zikwall/blogchain/actions"
	"github.com/zikwall/blogchain/middlewares"
)

func InitRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		Filter:           nil,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		ExposeHeaders:    nil,
		MaxAge:           0,
	}))

	app.Static("/docs", "./public/docs")
	app.Static("/uploads", "./public/uploads")

	// main endpoint group by `/api` prefix
	api := app.Group("/api", middlewares.JWT)
	// not usage JWT middleware in Login & Register endpoints
	auth := app.Group("/auth", middlewares.Auth)
	// editor required authorization
	editor := api.Group("/editor", middlewares.Authorization)

	v1 := api.Group("/v1")
	v1.Get("/profile/:username", actions.Profile)
	v1.Get("/content/:id", actions.GetContent)
	v1.Get("/contents/:page?", actions.GetContents)
	v1.Get("/tags", actions.Tags)
	v1.Get("/tag/:tag/:page?", actions.GetContents)
	auth.Post("/register", actions.Register)
	auth.Post("/login", actions.Login)
	auth.Post("/logout", actions.Login)

	editor.Get("/content/:id", actions.GetEditContent)
	editor.Post("/content/add", actions.AddContent)
	editor.Post("/content/update/:id", actions.UpdateContent)

	// todo deprecated delete
	api.Get("/", actions.HelloWorldAction)
}
