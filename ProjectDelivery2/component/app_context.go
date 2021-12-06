package component

import (
	"ProjectDelivery2/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetDbConnection() *gorm.DB
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	 db *gorm.DB
	 secretKey string
	 pb  pubsub.Pubsub
}
func NewAppContext(db *gorm.DB ,secretKey string ,pb pubsub.Pubsub) *appCtx{
	return &appCtx{db : db , secretKey: secretKey , pb:pb}
}

func(ctx *appCtx) GetDbConnection() *gorm.DB{
	return ctx.db
}
func(ctx *appCtx) SecretKey() string{
	return  ctx.secretKey
}
func(ctx *appCtx) GetPubSub() pubsub.Pubsub  {
	return ctx.pb
}