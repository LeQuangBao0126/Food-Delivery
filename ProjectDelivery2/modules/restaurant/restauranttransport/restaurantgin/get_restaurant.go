package restaurantgin

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/modules/restaurant/restaurantbiz"
	"ProjectDelivery2/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
	     uid,err := common.FromBase58(c.Param("id"))

		//go func(){
		//	defer common.AppRecover()
		//	panic("err dau tien")
		//}()

		//id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest , err)
			return
		}

		store := restaurantstorage.NewSqlStorage(appCtx.GetDbConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(),int(uid.GetLocalID()))

		if err != nil {
			 c.JSON(http.StatusBadRequest , err)
			 return
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.NewSimpleSuccessRes(data))
	}
}