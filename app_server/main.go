package main

import (
	"io/ioutil"
	"log"
	"os"

	"net/http"
	"github.com/labstack/echo/v4"
)

func account(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, it's account API")
}

func main() {
	logger := get()
	logger.trace("Starting the application...")
	logger.info("Something noteworthy happened...")
	logger.warn("There is something you should know about...")
	logger.error("Something went wrong...")

	e := echo.New()
	e.GET("/account", account)

	e.Start(":8000")
	

}

type Logger struct {
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

func (l *Logger) trace(message string) {
	l.Trace.Println(message)
}

func (l *Logger) info(message string) {
	l.Info.Println(message)
}

func (l *Logger) warn(message string) {
	l.Warn.Println(message)
}

func (l *Logger) error(message string) {
	l.Error.Println(message)
}

func get() Logger {
	l := Logger{
		log.New(ioutil.Discard, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return l
}