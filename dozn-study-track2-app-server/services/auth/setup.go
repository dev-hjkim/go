package auth

import (
	"dozn/app-server/logging"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	logging.Info("Setup auth router...")

	router.Post("/register", register)
	router.Post("/login", login)
	router.Post("/verify", verify)
}
