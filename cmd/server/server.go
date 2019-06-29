package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
    "github.com/galileo-pkm/go-mod-test/model"
)

func main() {
	test := model.CU()
_ = test	
e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}

