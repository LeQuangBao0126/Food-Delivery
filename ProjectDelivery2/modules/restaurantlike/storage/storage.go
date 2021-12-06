package restaurantlikestorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStorage (db *gorm.DB) *sqlStore{
	return &sqlStore{db : db}
}