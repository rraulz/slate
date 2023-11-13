package main

import (
	"context"
	"fmt"
	"os"
	"slate/api/controller"
	"slate/api/database"
	"slate/templates/login"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	pgInstance, err := database.NewPG(ctx, "host=localhost dbname=postgres user=postgres password=postgres")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer pgInstance.Close()

	e := echo.New()
	e.Static("/assets", "assets")

	// Routes
	e.GET("/", func(c echo.Context) error {
		component := login.Page()
		return component.Render(ctx, c.Response().Writer)
	})

	e.POST("/login", controller.LoginHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
