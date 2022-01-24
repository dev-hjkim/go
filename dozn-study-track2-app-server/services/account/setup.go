package account

import (
	"dozn/app-server/logging"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	logging.Info("Setup account router...")

	router.Post("/create", create)
	router.Post("/list", list)
}
