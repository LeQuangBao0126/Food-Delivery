package restaurantmodel

import (
	"ProjectDelivery2/common"
 )

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
	//logo
	//cover
	UserId     int             `json:"user_id" gorm:"column:user_id;"`
	User       *common.SimpleUser   `json:"user" `
	LikedCount int             `json:"liked_count" gorm:"column:liked_count"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (r *RestaurantUpdate) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"addr" gorm:"column:addr;"`
	UserId int    `json:"user_id" gorm:"column:user_id;"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

func (data *Restaurant) Mask(isAdminOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
