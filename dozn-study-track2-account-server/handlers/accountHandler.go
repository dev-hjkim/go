package handlers

import (
	"dozn/account-server/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAccountListHandler(c *fiber.Ctx) error {
	fmt.Println("GetAccountListHandler started!!")

	acct := models.Account{15, "01", 0.05}

	// JSON 인코딩
	jsonBytes, err := json.Marshal(acct)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	return c.SendString(jsonString)
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
