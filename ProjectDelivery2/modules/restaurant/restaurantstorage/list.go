package restaurantstorage

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/restaurant/restaurantmodel"
	"context"
)

func(s *sqlStore) ListDataByCondition(
	context context.Context ,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter ,
	paging *common.Paging ,
	moreKeys ...string ) ([]restaurantmodel.Restaurant,error) {

	var result []restaurantmodel.Restaurant

	db  := s.db.Table("restaurants")



	db = db.Where(conditions)

	if v:= filter ; v!= nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ? ",v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error ;err != nil{
		return nil, common.ErrDB(err)
	}
	for i:= range moreKeys{
		db = db.Preload(moreKeys[i])
	}
	if paging.FakeCursor != ""{
		uid ,_ := common.FromBase58(paging.FakeCursor)
		db = db.Where("id <  ? ", int(uid.GetLocalID()))
	}else{
		db = db.Offset((paging.Page- 1)* paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error ;err!= nil{
		return nil,err
	}

	return result, nil
}