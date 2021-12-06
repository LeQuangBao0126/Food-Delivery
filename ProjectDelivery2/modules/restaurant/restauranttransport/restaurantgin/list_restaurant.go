package restaurantgin

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/modules/restaurant/restaurantbiz"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"ProjectDelivery2/modules/restaurant/restaurantrepo"
	"ProjectDelivery2/modules/restaurant/restaurantstorage"
	restaurantlikestorage "ProjectDelivery2/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
)

func ListRestaurant(ctx component.AppContext)  gin.HandlerFunc{
	return func (c *gin.Context) {
		var filter restaurantmodel.Filter
		var paging common.Paging

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON( 400 , gin.H{
				"err":err.Error(),
			})
			return
		}
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON( 400 , gin.H{
				"err":err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStorage(ctx.GetDbConnection())
		likeStore := restaurantlikestorage.NewSqlStorage(ctx.GetDbConnection())
		// can be use with redis or mongo
		repo := restaurantrepo.NewRestaurantRepo(store,likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		results,err := biz.ListRestaurant(c.Request.Context() ,&filter, &paging)
		if err!= nil{
		    	c.JSON(400, gin.H{"err" : err.Error() })
				return
		}
		//truoc khi di về client => mask

		for i  := range results{
			results[i].Mask(false)

			if i == len(results) -1 {
				paging.NextCursor = results[i].FakeId.String()
			}
		}
		c.JSON(200, common.NewSuccessRes(results,paging,filter))
		return
	}
}










