package restaurantstorage

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

// Create ...
func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	return s.db.Create(data).Error
}
