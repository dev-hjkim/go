package transaction

import "github.com/gofiber/fiber/v2"

func list(context *fiber.Ctx) error {
	return context.SendString("TRANSACTION : list")
}

func deposit(context *fiber.Ctx) error {
	return context.SendString("TRANSACTION : deposit")
}

func transfer(context *fiber.Ctx) error {
	return context.SendString("TRANSACTION : transfer")
}
