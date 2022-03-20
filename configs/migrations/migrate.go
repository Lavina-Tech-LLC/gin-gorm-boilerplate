package main

import (
	"lvn-tools/configs"
	"lvn-tools/models"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
)

func init() {
	configs.StartViper()
}

// running go run configs/migrations/migrate.go
func main() {
	configs.Connect()
	configs.GetDB.AutoMigrate(&models.Logs{})
	llog.Info("Success")
}
