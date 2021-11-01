package transaction

import (
	"dozn/app-server/logging"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	logging.Info("Setup transaction router...")

	router.Post("/list", list)
	router.Post("/deposit", deposit)
	router.Post("/transfer", transfer)
}
