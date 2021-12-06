package ginrestaurantlike

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	rslikebiz "ProjectDelivery2/modules/restaurantlike/biz"
	restaurantlikemodel "ProjectDelivery2/modules/restaurantlike/model"
	restaurantlikestorage "ProjectDelivery2/modules/restaurantlike/storage"

	"github.com/gin-gonic/gin"
)

// /v1/restaurants/:id/liked_users
func ListUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		var paging common.Paging
		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
		}

		store := restaurantlikestorage.NewSqlStorage(ctx.GetDbConnection())
		// can be use with redis or mongo
		biz := rslikebiz.NewListUserLikeRestaurant(store)

		results, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(400, gin.H{"err": err.Error()})
			return
		}
		//truoc khi di về client => mask
		//for i  := range results{
		//	results[i].Mask(false)
		//
		//	if i == len(results) -1 {
		//		paging.NextCursor = results[i].FakeId.String()
		//	}
		//}

		c.JSON(200, common.NewSuccessRes(results, paging, filter))
		return
	}
}
