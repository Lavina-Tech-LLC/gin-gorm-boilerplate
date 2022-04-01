package migrate

import (
	"gin-gorm-boilerplate/internal/dbCon"
	"go/importer"
	"reflect"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
)

// running go run internal/dbCon/migrations/migrate.go
func Main() {
	dbCon.Connect()
	pkg, _ := importer.Default().Import("gin-gorm-boilerplate/internal/models")

	for _, declName := range pkg.Scope().Names() {
		model := reflect.Zero(reflect.TypeOf(declName))
		dbCon.DB.AutoMigrate(&model)
	}

	llog.Info("Success")
}
