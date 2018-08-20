package helloworld

import (
	"github.com/labstack/echo"
	"net/http"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
