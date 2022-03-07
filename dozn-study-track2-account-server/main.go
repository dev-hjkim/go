package main

import (
	"dozn/account-server/handlers"
	"dozn/account-server/models"
	"dozn/account-server/database"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "dozn/account-server/docs"
)

/*
  Account Server
*/

// @title Account API
// @version 1.0
// @description This is a description for Account Api
// @contact.name API Support
// @contact.email hjkim@dozn.co.kr
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /
func main() {
	config := LoadConfigration("local")
	app := fiber.New()

	database.ConnectDB(config)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/list", handlers.GetAccountListHandler)
	app.Post("/create", handlers.PostAccountCreateHandler)

	app.Listen(":" + config.Port)
}

func LoadConfigration(env string) models.Config {
	var config models.Config
	file, err := os.Open("./config/" + env + "_config.json")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(file)

	jsonParser.Decode(&config)
	return config
}
