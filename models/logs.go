package models

import (
	"lvn-tools/configs"
)

type (
	LogsModels struct{}
)

func (log *Logs) Log() error {
	db := configs.GetDB
	err := db.Create(&log).Error
	return err
}
