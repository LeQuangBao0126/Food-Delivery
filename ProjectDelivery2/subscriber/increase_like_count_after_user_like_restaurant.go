package subscriber

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/modules/restaurant/restaurantstorage"
	"ProjectDelivery2/pubsub"
	"ProjectDelivery2/skio"
	"context"
)
type HasRestaurantId interface {
	GetRestaurantId() int
	GetRestaurantOwnerId () int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext,ctx context.Context){
		c ,_ := appCtx.GetPubSub().Subscribe(ctx,common.TopicUserLikeRestaurant)
		store:= restaurantstorage.NewSqlStorage(appCtx.GetDbConnection())
		go func() {
			for{
				 msg:= <-c
				 restaurant := (msg.Data()).(HasRestaurantId)
				_ = store.IncreaseLikeCount(ctx , restaurant.GetRestaurantId() )
			}
		}()
}
//Chỗ này
// I wish i cound do something like that
//func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) func(ctx context.Context, message *pubsub.Message)error {
//	store:= restaurantstorage.NewSqlStorage(appCtx.GetDbConnection())
//	return func(ctx context.Context , message *pubsub.Message)error{
//		restaurant:= (message.Data()).(HasRestaurantId)
//		return store.IncreaseLikeCount(ctx , restaurant.GetRestaurantId() )
//	}
//}
func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext ,
	rtEngine skio.RealtimeEngine) consumerJob{
	store:= restaurantstorage.NewSqlStorage(appCtx.GetDbConnection())
	return consumerJob{
		Title :"Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error{
			likeData := (message.Data()).(HasRestaurantId)

			rtEngine.EmitToUser(likeData.GetRestaurantOwnerId(),"TopicUserLikeRestaurant",likeData)

			return store.IncreaseLikeCount(ctx,likeData.GetRestaurantId())
		},
	}
}