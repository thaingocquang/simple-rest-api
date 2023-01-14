package restaurantmodel

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

func (RestaurantCreate) TableName() string {
	return "restaurants"
}
