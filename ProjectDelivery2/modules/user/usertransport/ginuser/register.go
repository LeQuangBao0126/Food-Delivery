package ginuser

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	 "ProjectDelivery2/component/hasher"
	"ProjectDelivery2/modules/user/userbiz"
	"ProjectDelivery2/modules/user/usermodel"
	"ProjectDelivery2/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context){
		db := appCtx.GetDbConnection()
		var data usermodel.UserCreate

		if err:= c.ShouldBind(&data); err != nil{
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}

		store:= userstorage.NewSqlStorage(db)
		hashers := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store,hashers)

		if err := biz.Register(c.Request.Context(),&data);err!= nil{
			c.JSON(400, common.ErrCannotCreateEntity("User",err))
			return
		}
	  //data.GenUID(common.DbTypeUser)
	  c.JSON(200,gin.H{"data":data.Id})
	}
}
