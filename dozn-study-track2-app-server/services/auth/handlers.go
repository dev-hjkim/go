package auth

import "github.com/gofiber/fiber/v2"

func register(context *fiber.Ctx) error {
	return context.SendString("AUTH : register")
}

func login(context *fiber.Ctx) error {
	return context.SendString("AUTH : login")
}

func verify(context *fiber.Ctx) error {
	return context.SendString("AUTH : verify")
}
