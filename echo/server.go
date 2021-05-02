package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Hello struct {
	Message string `json:"message"`
}

type Health struct {
	Status string `json:"status"`
}

func main() {
	e := echo.New()
	e.GET("/hello-world", handleHello)
	e.GET("/health", handleHealth)
	e.Logger.Fatal(e.Start(":9000"))
}

func handleHello(c echo.Context) error {
	message := &Hello{
		Message: "hello world",
	}
	return c.JSON(http.StatusOK, message)
}

func handleHealth(c echo.Context) error {
	status := &Health{
		Status: "healthy",
	}
	return c.JSON(http.StatusOK, status)
}
