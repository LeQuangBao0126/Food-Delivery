package rslikebiz

import (
	"ProjectDelivery2/component/asyncjob"
	restaurantlikemodel "ProjectDelivery2/modules/restaurantlike/model"
	"context"
)

type UserUnLikeRestaurantStore interface{
	Delete( context context.Context ,userId,restaurantId int) error

}
type decreaseStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}
type userUnLikeRestaurantBiz struct{
	store UserUnLikeRestaurantStore
	descStore decreaseStore
}

func NewUserUnLikeRestaurantBiz (store UserUnLikeRestaurantStore ,descStore decreaseStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store:store,
		descStore: descStore,
	}
}

func(biz *userUnLikeRestaurantBiz)  UnLikeRestaurant(ctx context.Context , userId,restaurantId int ) error {
	if err := biz.store.Delete(ctx, userId , restaurantId);err != nil{
		return  restaurantlikemodel.ErrCanNotUnLikeRestaurant(err)
	}
	//side effect

	job:= asyncjob.NewJob(func(ctx context.Context) error {
		return  biz.descStore.DecreaseLikeCount(ctx,restaurantId)
	})
	jobManager := asyncjob.NewGroup(false,job)

	jobManager.Run(ctx)
	return nil
}
