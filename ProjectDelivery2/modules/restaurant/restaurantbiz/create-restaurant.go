package restaurantbiz

import (
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"context"
)
type CreateRestaurantStore interface {
	 Create(context context.Context , data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBiz struct {
	 store  CreateRestaurantStore
}

func NewRestaurantBiz ( store CreateRestaurantStore ) *createRestaurantBiz{
	return &createRestaurantBiz { store : store }
}
func(  biz *createRestaurantBiz ) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate ) error {
	if err := biz.store.Create(ctx,data) ; err!= nil{
		return err
	}
	return nil
}