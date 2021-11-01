package handlers

import (
	"dozn/account-server/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAccountListHandler(c *fiber.Ctx) error {
	fmt.Println("GetAccountListHandler started!!")
	return c.SendStatus(fiber.StatusOK)
}

func PostAccountCreateHandler(c *fiber.Ctx) error {
	fmt.Println("PostAccountCreateHandler started!!")

	account := new(models.Account)

	if err := c.BodyParser(account); err != nil {
		return err
	}

	fmt.Println(account)

	return c.SendStatus(fiber.StatusOK)
}
