package migrate

import (
	"gin-gorm-boilerplate/internal/dbCon"
	"gin-gorm-boilerplate/internal/models"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
)

// running go run internal/dbCon/migrations/migrate.go
func Main() {	
	dbCon.ConnectPG()
	dbCon.GetDB.AutoMigrate(&models.Logs{})
	dbCon.GetDB.AutoMigrate(&models.Users{})
	llog.Info("Success")
}
