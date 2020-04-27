package main

import (
	"github.com/azer/go-makefile-example/src/hello-world"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"os"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.GET("/", helloworld.HelloWorld)
	e.Start(os.Getenv("ADDR"))
}
