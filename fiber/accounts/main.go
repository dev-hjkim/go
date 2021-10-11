package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go/fiber/accounts/handlers"
	"go/fiber/accounts/models"
	"fmt"
	"os"
	"encoding/json"
)

/*
  Account Server
*/
func main() {
	config := LoadConfigration("local")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Get("/all", handlers.GetAccountListHandler)
	app.Post("/create", handlers.PostAccountCreateHandler)

	app.Listen(":"+ config.Port)
}

func LoadConfigration(env string) models.Config {
    var config models.Config
    file, err := os.Open("./config/"+ env + "_config.json")
    defer file.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(file)
    jsonParser.Decode(&config)
    return config
}