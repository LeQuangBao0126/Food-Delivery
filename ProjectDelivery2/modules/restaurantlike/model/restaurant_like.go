package restaurantlikemodel

import (
	"ProjectDelivery2/common"
	"time"
)

type Like struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId     int 		 `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time    `json:"created_at" gorm:"column:created_at"`
	User *common.SimpleUser  `json:"user"`
}

func ( Like ) TableName () string {
	return "restaurant_likes"
}
func ( l *Like ) GetRestaurantId () int {
	return l.RestaurantId
}
func ( l *Like ) GetRestaurantOwnerId () int {
	return l.UserId
}

func ErrCanNotLikeRestaurant (err error  ) *common.AppError{
	return common.NewCustomError(err,"Cannot create restaurantLike","abc")
}
func ErrCanNotUnLikeRestaurant (err error  ) *common.AppError{
	return common.NewCustomError(err,"Cannot delete restaurantLike","abc")
}