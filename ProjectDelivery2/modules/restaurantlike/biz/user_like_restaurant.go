package rslikebiz

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/restaurantlike/model"
	"ProjectDelivery2/pubsub"
	"context"
)

type UserLikeRestaurantStore interface{
	Create( context context.Context , data *restaurantlikemodel.Like) error

}
type increaseStore interface{
	IncreaseLikeCount(ctx context.Context, id int) error
}
type userLikeRestaurantBiz struct{
	store UserLikeRestaurantStore
	//incStore increaseStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz (store UserLikeRestaurantStore,
	pubsub pubsub.Pubsub) *userLikeRestaurantBiz{
	return &userLikeRestaurantBiz{
		store:store,
		pubsub:pubsub,
	}
}

func(biz *userLikeRestaurantBiz)  LikeRestaurant(ctx context.Context ,
	data *restaurantlikemodel.Like) error {
	 if err := biz.store.Create(ctx,data);err != nil{
		 return  restaurantlikemodel.ErrCanNotLikeRestaurant(err)
	 }
	//side effect
	//job:= asyncjob.NewJob(func(ctx context.Context) error {
	//	return  biz.incStore.IncreaseLikeCount(ctx,data.RestaurantId)
	//})
	//asyncjob.NewGroup(false,job).Run(ctx)
	//cách mới sẽ là publish event

	 biz.pubsub.Publish(ctx,common.TopicUserLikeRestaurant,pubsub.NewMessage(data))

	 return nil
}
