package ginrestaurantlike

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/modules/restaurant/restaurantstorage"
	rslikebiz "ProjectDelivery2/modules/restaurantlike/biz"
	restaurantlikestorage "ProjectDelivery2/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
)

//POST /restaurants/:id/un-like
func UnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {
		uid , err := common.FromBase58(c.Param("id"))
		if err != nil{
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet("user").(common.Requester)

		store :=  restaurantlikestorage.NewSqlStorage(appCtx.GetDbConnection())
		unlikeStore := restaurantstorage.NewSqlStorage(appCtx.GetDbConnection())
		biz := rslikebiz.NewUserUnLikeRestaurantBiz(store,unlikeStore)

		if err:= biz.UnLikeRestaurant(c.Request.Context(),requester.GetUserId(),int(uid.GetLocalID())) ; err != nil {
			panic(err)
		}
		c.JSON(200,common.NewSimpleSuccessRes(true))
	}
}