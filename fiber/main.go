package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
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

	app.Group("/accounts/*", func(c *fiber.Ctx) error {
		doRequest("http://localhost:8001/" + c.Params("*"), c.Method(), c.Body())
		
		return c.Next()
	})

	app.Listen(":8000")
}

func doRequest(url string, method string, body []byte) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// request
	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")

	req.SetBody(body)

	fasthttp.Do(req, resp)

	// response
	bodyBytes := resp.Body()
	fmt.Println(string(bodyBytes))
}