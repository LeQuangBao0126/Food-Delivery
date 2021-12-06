package userstorage

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/modules/user/usermodel"
	"context"
	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	_ , span := trace.StartSpan(ctx,"user.storage.findUser")
	defer span.End()
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}