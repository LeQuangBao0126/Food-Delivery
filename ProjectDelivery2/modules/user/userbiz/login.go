package userbiz

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"ProjectDelivery2/component/tokenprovider"
	"ProjectDelivery2/modules/user/usermodel"
	"context"
	"go.opencensus.io/trace"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	ctx1 , span := trace.StartSpan(ctx,"user.biz.login")
	span.AddAttributes(  trace.Int64Attribute("user-id", 1 ) )

	user, err := business.storeUser.FindUser(ctx1, map[string]interface{}{"email": data.Email})

	span.End()


	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	_ , span2 := trace.StartSpan(ctx,"user.biz.gen-jwt")
	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	span2.End()

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := business.tokenProvider.Generate(payload, business.expiry * 7)
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	//account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
	//return account, nil
}