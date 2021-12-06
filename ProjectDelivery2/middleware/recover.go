package middleware

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"github.com/gin-gonic/gin"
)

func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					//panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				//panic(err)
				return
			}
		}()

		c.Next()
	}
}