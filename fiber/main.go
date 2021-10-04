package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
	"encoding/json"
)


/*
  App Server
*/
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Group("/accounts", func(c *fiber.Ctx) error {
		doRequest("http://localhost:8001" + c.Path(), c.Method(), string(c.Body()[:]))
		
		var data []byte
    	json.Unmarshal(c.Body(), &data)
		fmt.Println(data)
		
		return c.Next()
	})
	//Account.Get("/list", service.AccountHandler)

	app.Listen(":8000")
}

func doRequest(url string, method string, body string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.SetBodyString(body)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	fmt.Println(string(bodyBytes))
	// User-Agent: fasthttp
	// Body:
}