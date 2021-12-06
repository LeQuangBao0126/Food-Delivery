package middleware

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/component/jwt"
	"ProjectDelivery2/modules/user/userstorage"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

//  1.Get token from header
//  2.Validate token adn parse from payload
// 	3.From token payload use user from its request
func extraTokenFromHeaderString( token string ) (string,error){
	tokens := strings.Split( token, " ")

	if tokens[0] != "Bearer " && tokens[1] == ""{
		return "" , errors.New("wrong authentication token")
	}
	return tokens[1] , nil
}
func RequireAuth(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		token ,err := extraTokenFromHeaderString(c.Request.Header.Get("Authorization"))

		db:= appCtx.GetDbConnection()
		store:= userstorage.NewSqlStorage(db)


		 payload,err:= tokenProvider.Validate(token);
		if err != nil{
			c.AbortWithStatusJSON(400,err)
			return
		}

		user,err := store.FindUser(c.Request.Context() ,
			map[string]interface{}{"id": payload.UserId })

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("User no permission")))
		}

		c.Set("user",user)
		c.Next()

	}
}
