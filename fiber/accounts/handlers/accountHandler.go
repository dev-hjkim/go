package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go/fiber/accounts/model"
)

func ListHandler(c *fiber.Ctx) error {
	fmt.Println("ListHandler started!!")
	return c.SendStatus(fiber.StatusOK)
}

func AccountHandler(c *fiber.Ctx) error {
	fmt.Println("AccountHandler started!!")

	a := new(model.Account)

	fmt.Println(c.Body())

	if err := c.BodyParser(a); err != nil {
		return err
	}

	fmt.Println(a.AccoutSeq)
	fmt.Println(a.AccountName)

	return c.SendStatus(fiber.StatusOK)
}