package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go/fiber/accounts/handlers"
)

/*
  Account Server
*/
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Get("/accounts/list", handlers.ListHandler)
	app.Post("/accounts/account", handlers.AccountHandler)

	app.Listen(":8001")
}