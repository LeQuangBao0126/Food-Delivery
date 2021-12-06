package ginuser

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/component/hasher"
	"ProjectDelivery2/component/jwt"
	"ProjectDelivery2/modules/user/userbiz"
	"ProjectDelivery2/modules/user/usermodel"
	"ProjectDelivery2/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDbConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSqlStorage(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessRes(account))
	}
}