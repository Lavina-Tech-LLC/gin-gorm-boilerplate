package models

import "gin-gorm-boilerplate/internal/dbCon"

type (
	UsersModel struct{}
)

func (usr *Users) GetUserByKey() error {

	res := dbCon.DB.First(&usr).Where("key = ?", usr.Key)

	return res.Error
}
