package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"hello-world"
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
