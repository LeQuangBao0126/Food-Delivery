package restaurantgin

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/modules/restaurant/restaurantbiz"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"ProjectDelivery2/modules/restaurant/restaurantstorage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		 uid,_ := common.FromBase58(c.Param("id"))

		// id, _ := strconv.Atoi(c.Param("id"))
	    	var a  restaurantmodel.RestaurantUpdate

		 if err := c.ShouldBind(&a); err != nil {
			c.JSON( 400 , gin.H{
				"err":err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSqlStorage(appCtx.GetDbConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		fmt.Println(int(uid.GetLocalID()))
		if err := biz.UpdateRestaurant(c.Request.Context(),int(uid.GetLocalID()), &a);err != nil{
			c.JSON(400,gin.H{"err":err})
			return
		}
		c.JSON(http.StatusOK, common.NewSimpleSuccessRes(a))
	}
}
