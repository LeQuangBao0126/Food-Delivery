package restaurantrepo

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"context"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		context context.Context ,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter ,
		paging *common.Paging ,
		moreKeys ...string ) ([]restaurantmodel.Restaurant,error)
}
type LikeStore interface {
	GetRestaurantLikes(context context.Context,ids []int)(map[int]int , error)
}
type listRestaurantRepo struct {
	store  ListRestaurantStore
	likeStore LikeStore
}
func NewRestaurantRepo ( store  ListRestaurantStore , likeStore LikeStore) *listRestaurantRepo{
	return &listRestaurantRepo{ store: store,likeStore: likeStore}
}

func(  repo *listRestaurantRepo) ListRestaurantRepo(
	ctx context.Context,
	filter *restaurantmodel.Filter ,
	paging *common.Paging ,
) ([]restaurantmodel.Restaurant,error) {

	results , err := repo.store.ListDataByCondition(ctx,nil,filter,paging,"User")
	if err!= nil {
		return nil,err
	}

	//ids := make([]int , len(results))
	//for i := range results{
	//	ids[i] = results[i].Id
	//}
	//
	//mapResLike , err := repo.likeStore.GetRestaurantLikes(ctx,ids)
	//
	//if err!= nil{
	//	//co the bo qua dữ lieu like_count để cứu lấy data restaurant
	//}
	//if v := mapResLike ; v!= nil {
	//	for i := range results{
	//		results[i].LikedCount = v[results[i].Id]
	//	}
	//}

	return results ,nil
}