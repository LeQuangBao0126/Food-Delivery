package ginuser

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {
		data := c.MustGet("user").(common.Requester)

		//neu data ok
		c.JSON(200,gin.H{ "data":data })
	}
}