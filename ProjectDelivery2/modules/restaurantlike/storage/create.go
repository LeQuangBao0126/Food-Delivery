package restaurantlikestorage

import (
	"ProjectDelivery2/common"
	 "ProjectDelivery2/modules/restaurantlike/model"
	"context"
)

func(s *sqlStore) Create( context context.Context , data *restaurantlikemodel.Like) error{
	db  := s.db

	if err := db.Table(data.TableName()).Create(data).Error ; err !=nil {
		return common.ErrDB(err)
	}
	return nil
}