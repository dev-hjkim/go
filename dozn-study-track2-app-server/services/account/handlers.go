package account

import (
	"github.com/gofiber/fiber/v2"
)

func create(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : create")
}

func list(context *fiber.Ctx) error {
	res := doRequest("http://localhost:3001/list", context.Method(), context.Body())

	return context.SendString(res)
}
