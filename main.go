package main

import (
	"github.com/cumedang/GoFile/sanz"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

var a string

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleSanz(c echo.Context) error {
	si := c.FormValue("si")
	sanz.Select(si)
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/sanz", handleSanz)
	e.Logger.Fatal(e.Start(":2234"))
}
