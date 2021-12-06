package restaurantstorage

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
		//nếu user ở 1 service khác database khác thì if ở đây và call API rest
	}

	if err := db.Where(conditions).First(&result).Error; err != nil {
		 	if err == gorm.ErrRecordNotFound{
				return nil, common.RecordNotFound
			}
			return nil,common.ErrDB(err)
	}

	return &result, nil
}
