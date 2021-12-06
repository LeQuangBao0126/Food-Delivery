package restaurantstorage

import (
	"ProjectDelivery2/common"
	"context"
)

func(s *sqlStore) Delete( context context.Context ,id int ) error{
	db  := s.db.Table("restaurants")
	if err := db.Where("id= ? ",id).Delete(nil).Error ; err !=nil {
		return common.ErrDB(err)
	}
	return nil
}