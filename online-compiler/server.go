package main

import (
	"online-compiler/execute"
	"online-compiler/hello"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", hello.Hello)
    e.POST("/execute", execute.RunGoCode)
	e.Logger.Fatal(e.Start(":8080"))
}
