package services

import (
	"atlantis/dbs"
	"atlantis/model"
)

//UserCreate 添加用户
func UserCreate(username, password, createdBy string) error {
	user := model.User{
		Username:  username,
		Password:  password,
		CreatedBy: createdBy,
	}
	return dbs.DB.Create(&user).Error
}

//UserDetail 用户详情
func UserDetail(id int) (result model.User, err error) {
	err = dbs.DB.Where("id = ? AND is_deleted = 0").First(&result).Error
	return result, err
}
