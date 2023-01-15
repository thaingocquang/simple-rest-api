package restaurantmodel

import (
	"errors"
	"strings"
)

type Restaurant struct {
	ID      int    `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	Address string `gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	ID      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:addr"`
}

func (rc RestaurantCreate) Validate() error {
	rc.Name = strings.TrimSpace(rc.Name)

	if len(rc.Name) == 0 {
		return errors.New("restaurant name can not be blank")
	}

	return nil
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}
