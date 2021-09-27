package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"

	"account/repositories"
	"account/usecases"
	"account/handlers"
)

type (
	Account struct {
		AccountId string `json:"accountId" validate:"required"`
		AccountNum string `json:"accountNum" validate:"required"`
		Amount int `json:"amount"`	// 없으니  default인 0으로
	}

	CustomValidator struct {
		validtr *validator.Validate	// validator -> validtr
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validtr.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}



func main() {
	// https 서버 세팅하는 법
	e := echo.New()
	// if err := e.StartTLS(":8443", "server.crt", "server.key"); err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }

	// structure setting
	// accountRepository := repositories.NewAccountRepository()
	// accountUsecase := usecases.NewAccountUsecase(accountRepository)
	// accountHandler := handlers.NewAccountHandler(accountUsecase)

	// 미들웨어 - cors 설정
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  }))

	// get request 처리
	e.GET("/accounts/:number", func(c echo.Context) error {
		number := c.Param("number")
		return c.String(http.StatusOK, number)
	})

	// // post request 처리
	e.Validator = &CustomValidator{validtr: validator.New()}
	e.POST("/account", func(c echo.Context) (err error) {
		u := new(Account)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, u)
	})

	e.Start(":8080")
}
