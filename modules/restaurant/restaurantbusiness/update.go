package restaurantbusiness

import (
	"context"
	"errors"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBusiness struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBusiness(store UpdateRestaurantStore) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (business *updateRestaurantBusiness) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	oldData, err := business.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err = business.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
