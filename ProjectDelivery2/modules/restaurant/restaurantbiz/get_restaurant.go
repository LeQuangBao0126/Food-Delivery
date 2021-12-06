package restaurantbiz

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"context"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {

	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id},"User")
	if err != nil {
		if err == common.RecordNotFound{
			return nil , common.ErrEntityNotFound("restaurant",err)
		}
		return nil,err
	}
	return data, err
}