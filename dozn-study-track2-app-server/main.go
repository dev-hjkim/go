package main

import (
	"dozn/app-server/logging"
	"dozn/app-server/services/account"
	"dozn/app-server/services/auth"
	"dozn/app-server/services/transaction"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	logging.Info("Preparing server...")

	// Echo instance
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Routes
	logging.Info("Setup route configurations...")
	api := app.Group("/api")

	// 1. Registration (User)
	// 2. Login (User)
	// 3. Auto login (User)
	// 4. Create bank account (User, Account)
	// 5. Get bank accounts (User, Account)
	// 6. Transaction list for one bank account (User, Account, Transaction)
	// 7. Deposit (User, Account, Transaction)
	// 8. Money tranfer (User, Account, Transaction)

	auth.SetupRoutes(api.Group("/auth"))
	account.SetupRoutes(api.Group("/account"))
	transaction.SetupRoutes(api.Group("/transaction"))

	// Start server
	logging.Info("Start server...")
	log.Fatal(app.Listen(":3000"))
}
