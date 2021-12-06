package ginrestaurantlike

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	rslikebiz "ProjectDelivery2/modules/restaurantlike/biz"
	restaurantlikemodel "ProjectDelivery2/modules/restaurantlike/model"
	restaurantlikestorage "ProjectDelivery2/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
)

//POST /restaurants/:id/like
func LikeRestaurant(appCtx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {
		uid , err11 := common.FromBase58(c.Param("id"))
		if err11 != nil {
			panic(common.ErrInvalidRequest(err11))
		}

		var data restaurantlikemodel.Like



		requester := c.MustGet("user").(common.Requester)
		data.UserId = requester.GetUserId()
		data.RestaurantId = int(uid.GetLocalID())

		store := restaurantlikestorage.NewSqlStorage(appCtx.GetDbConnection())
		//likeStore := restaurantstorage.NewSqlStorage(appCtx.GetDbConnection())

		biz := rslikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub() )

		if err:= biz.LikeRestaurant(c.Request.Context(),&data) ; err != nil {
			 panic(err)
		}
		c.JSON(200,common.NewSimpleSuccessRes(true))
	}
}