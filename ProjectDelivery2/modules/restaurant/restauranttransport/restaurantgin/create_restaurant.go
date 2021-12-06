package restaurantgin

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/modules/restaurant/restaurantbiz"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"ProjectDelivery2/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
)

func CreateRestaurant(ctx component.AppContext)  gin.HandlerFunc{
	return func (c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON( 400 , gin.H{
				"err":err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSqlStorage(ctx.GetDbConnection())
		biz := restaurantbiz.NewRestaurantBiz(store)

		requester := c.MustGet("user").(common.Requester)
		data.UserId = requester.GetUserId()

		if err := biz.CreateRestaurant(c.Request.Context() , &data); err!= nil{
		    	c.JSON(400, gin.H{"err" : err.Error() })
				return
		}
		c.JSON(200, data)
		return
	}
}










