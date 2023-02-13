package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Where(condition).
		First(&result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, err
	}

	return &result, nil
}
