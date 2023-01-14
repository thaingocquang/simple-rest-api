package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-rest-api/component"
	"simple-rest-api/modules/restaurant/restauranttransport/echorestaurant"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/deliveryfood?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db, err)

	e := echo.New()

	appCtx := component.NewAppContext(db)

	restaurants := e.Group("/restaurants")

	restaurants.POST("", echorestaurant.CreateRestaurant(appCtx))
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	e.Logger.Fatal(e.Start(":1323"))
}
