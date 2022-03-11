package handlers

import (
	"dozn/account-server/models"
	"dozn/account-server/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAccountList godoc
// @Summary Get an account list
// @Description Get an account list by userId
// @ID get-account-list
// @Accept  json
// @Produce  json
// @Tags Account
// @Success 200 {object} models.Account
// @Router /list [get]
func GetAccountListHandler(c *fiber.Ctx) error {
	fmt.Println("GetAccountListHandler started!!")

	db := database.DB
	var accounts []models.Account

	db.Find(&accounts)

	// If no account is present return an error
    if len(accounts) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No accounts present", "data": nil})
    }

    // Else return accounts
    return c.JSON(fiber.Map{"status": "success", "message": "Accounts Found", "data": accounts})
}


// PostAccountCreate godoc
// @Summary Post an account
// @Description Post an account
// @ID create-account
// @Accept  json
// @Produce  json
// @Tags Account
// @Success 200 {object} models.Account
// @Router /create [post]
func PostAccountCreateHandler(c *fiber.Ctx) error {
	fmt.Println("PostAccountCreateHandler started!!")

	db := database.DB
	account := new(models.Account)

	if err := c.BodyParser(account); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create the Account and return error if encountered
    err := db.Create(&account).Error
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create account", "data": err})
    }

    // Return the created account
    return c.JSON(fiber.Map{"status": "success", "message": "Created Account", "data": account})
}
