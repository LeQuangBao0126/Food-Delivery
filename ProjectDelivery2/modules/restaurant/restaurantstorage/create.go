package restaurantstorage

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"context"
)

func(s *sqlStore) Create( context context.Context , data *restaurantmodel.RestaurantCreate) error{
	db  := s.db

	if err := db.Table(data.TableName()).Create(data).Error ; err !=nil {
		return common.ErrDB(err)
	}
	return nil
}