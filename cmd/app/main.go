package main

import (
	"log"

	"github.com/Giafn/goMoneySaveAndManage/configs"
	"github.com/Giafn/goMoneySaveAndManage/internal/postgres"
	"github.com/Giafn/goMoneySaveAndManage/internal/router"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadConfig()

	e := echo.New()

	postgres.InitDB()

	router.InitRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
