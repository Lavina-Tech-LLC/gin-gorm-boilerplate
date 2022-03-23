package models

import "gin-gorm-boilerplate/internal/dbCon"

type (
	LogsModels struct{}
)

func (log *Logs) Log() error {
	db := dbCon.GetDB
	err := db.Create(&log).Error
	return err
}
