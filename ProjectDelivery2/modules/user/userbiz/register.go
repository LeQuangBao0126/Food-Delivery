package userbiz

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/user/usermodel"
	"context"
)

type RegisterStorage interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type Hasher interface {
	Hash(data string) string
}
type registerBusiness struct{
	registerStorage  RegisterStorage
	hasher Hasher
}

func NewRegisterBusiness( registerStorage RegisterStorage , hasher Hasher ) *registerBusiness{
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher: hasher,
	}
}
func (business *registerBusiness) Register(ctx context.Context,data *usermodel.UserCreate) error{
	user , _ := business.registerStorage.FindUser(ctx,map[string]interface{}{"email": data.Email})
	if user != nil{
		return usermodel.ErrEmailExisted
	}

	salt:= common.GenSalt(50)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err:= business.registerStorage.CreateUser(ctx,data);err!= nil{
		return common.ErrCannotCreateEntity(usermodel.EntityName,err)
	}

	return nil
}