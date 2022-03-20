package configs

import (
	"flag"
	"fmt"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// GetDB variable for connection DB
	GetDB *gorm.DB // connection about DB
)

func Connect() {

	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		llog.Error(err)
	}

	GetDB = db
}

func StartViper() {
	runMode := flag.String("m", "dev", "runMode, by default it is dev")

	viper.SetConfigName(fmt.Sprintf("config_%s", *runMode))
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")

	err := viper.ReadInConfig()
	if err != nil {
		llog.Notice("Hello world")
	}
}
