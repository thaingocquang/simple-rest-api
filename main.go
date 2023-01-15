package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"simple-rest-api/component"
	"simple-rest-api/modules/restaurant/restauranttransport/echorestaurant"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/deliveryfood?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	e := echo.New()

	appCtx := component.NewAppContext(db)

	restaurants := e.Group("/restaurants")

	restaurants.POST("", echorestaurant.CreateRestaurant(appCtx))
	restaurants.GET("", echorestaurant.ListRestaurant(appCtx))

	e.Logger.Fatal(e.Start(":1323"))
}
