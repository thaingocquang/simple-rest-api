package restaurantbusiness

import (
	"context"
	"errors"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(
	ctx context.Context,
	id int,
) error {
	oldData, err := business.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err = business.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}

	return nil
}
