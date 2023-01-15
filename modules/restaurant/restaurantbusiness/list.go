package restaurantbusiness

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	store ListRestaurantStore
}

func NewListRestaurantBusiness(store ListRestaurantStore) *listRestaurantBusiness {
	return &listRestaurantBusiness{store: store}
}

func (business *listRestaurantBusiness) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	return business.store.ListDataByCondition(ctx, nil, filter, paging)
}
