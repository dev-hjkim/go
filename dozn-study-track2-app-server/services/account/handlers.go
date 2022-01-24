package account

import (
	"github.com/gofiber/fiber/v2"
)

func create(context *fiber.Ctx) error {
	res := doRequest("http://localhost:3001/create", context.Method(), context.Body())
	return context.SendString(res)
}

func list(context *fiber.Ctx) error {
	res := doRequest("http://localhost:3001/list", context.Method(), context.Body())

	return context.SendString(res)
}
