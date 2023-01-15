package restaurantstorage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition).Where("status in (1)")

	if filter != nil {
		if filter.CityID > 0 {
			db = db.Where("city_id = ?", filter.CityID)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
