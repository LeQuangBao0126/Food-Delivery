package restaurantlikestorage

import (
	"ProjectDelivery2/common"
	restaurantlikemodel "ProjectDelivery2/modules/restaurantlike/model"
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

func (s *sqlStore) GetRestaurantLikes(context context.Context, ids []int) (map[int]int, error) {

	results := make(map[int]int, len(ids))

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikedCount   int `gorm:"column:liked_count"`
	}
	var listLike []sqlData

	db := s.db

	if err := db.Table("restaurant_likes").Select("restaurant_id, count(restaurant_id) as liked_count ").
		Where("restaurant_id in (?) ", ids).
		Group("restaurant_id").
		Find(&listLike).Error; err != nil {
		return nil, err
	}

	for _, item := range listLike {
		results[item.RestaurantId] = item.LikedCount
	}

	return results, nil
}

func (s *sqlStore) GetUserLikesRestaurant(
	context context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {

	var result []restaurantlikemodel.Like

	db := s.db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ? ", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	db = db.Preload("User")

	if paging.FakeCursor != "" {
		uid, _ := common.FromBase58(paging.FakeCursor)
		db = db.Where("created_at <  ? ", int(uid.GetLocalID()))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).
		Error; err != nil {
		return nil, err
	}

	users := make([]common.SimpleUser, len(result))
	for i := range users {
		users[i] = *result[i].User
	}

	for i, item := range result {
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.
				Format("2006-01-02T15:04:05-0700"))))
			paging.NextCursor = cursorStr
		}
	}
	return users, nil
}
