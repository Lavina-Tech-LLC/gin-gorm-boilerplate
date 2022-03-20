package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Logs struct {
		Time  time.Time
		Level string
		Code  int //`gorm:"autoIncrement"`
		Msg   string
		gorm.Model
	}
)
