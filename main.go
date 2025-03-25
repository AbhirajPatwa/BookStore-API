package main

import (
	"net/http"
	"github.com/AbhirajPatwa/BookStore-API/config"
	"github.com/AbhirajPatwa/BookStore-API/controller"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})
	bookRoute := e.Group("/book")
	bookRoute.POST("/", controller.CreateBook)
	bookRoute.GET("/:id", controller.GetBook)
	bookRoute.PUT("/:id", controller.UpdateBook)
	bookRoute.DELETE("/:id", controller.DeleteBook)
	e.Logger.Fatal(e.Start(":8080"))

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()
	
	e.Logger.Fatal(e.Start(":8080"))
}