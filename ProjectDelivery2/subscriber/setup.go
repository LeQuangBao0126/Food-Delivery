package subscriber

import (
	"ProjectDelivery2/component"
	"context"
)

func SetupSubscribers(appCtx component.AppContext){
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx,context.Background())
}