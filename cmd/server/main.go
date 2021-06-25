package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apiv1 "ternary.app/backend-assignment/pkg/api/v1"
)

func main() {
	// Echo instance
	var e = echo.New()

	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiv1.Register(e)

	var err = e.Start(":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
