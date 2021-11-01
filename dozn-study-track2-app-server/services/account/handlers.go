package account

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func create(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : create")
}

func List(context *fiber.Ctx) error {
	fmt.Println("http://localhost:3001/"+context.Params("*"))
	doRequest("http://localhost:3001/"+context.Params("*"), context.Method(), context.Body())

	return context.Next()
}