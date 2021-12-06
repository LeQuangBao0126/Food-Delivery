package rslikebiz

import (
	"ProjectDelivery2/common"
	restaurantlikemodel "ProjectDelivery2/modules/restaurantlike/model"
	"context"
)

type ListUserLikeRestaurantStore interface{
	GetUserLikesRestaurant(
		context context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurant(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz{
	return &listUserLikeRestaurantBiz{
		store: store,
	}
}
func (biz *listUserLikeRestaurantBiz) ListUsers(ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging) ([]common.SimpleUser,error) {

	users ,err := biz.store.GetUserLikesRestaurant(ctx,nil,filter,paging)
	if err != nil{
		return nil, common.ErrCannotListEntity("user",err)
	}

	return users,nil
}
