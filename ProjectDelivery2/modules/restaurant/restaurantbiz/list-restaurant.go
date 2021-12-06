package restaurantbiz


import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"context"
)

type ListRestaurantRepo interface {
	ListRestaurantRepo(
		ctx context.Context,
		filter *restaurantmodel.Filter ,
		paging *common.Paging ,
	) ([]restaurantmodel.Restaurant,error)
}
type listRestaurantBiz struct {
	 repo ListRestaurantRepo
}

func NewListRestaurantBiz ( repo ListRestaurantRepo ) *listRestaurantBiz {
	return &listRestaurantBiz{ repo: repo}
}

func(  biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter ,
	paging *common.Paging ,
) ([]restaurantmodel.Restaurant,error) {

	results , err :=biz.repo.ListRestaurantRepo(ctx,filter,paging)
	if err != nil{
		return nil , err
	}
	return results ,nil
}