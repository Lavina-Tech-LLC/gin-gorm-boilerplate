package dbCon

import (
	"database/sql"
	"fmt"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
	"github.com/casbin/casbin/v2"
	casbinpgadapter "github.com/cychiuae/casbin-pg-adapter"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// GetDB variable for connection DB
	RunMode    = ""
	DB         *gorm.DB // connection about DB
	Casbin     *casbin.Enforcer
	ConfigPath = ""
	Dm         = false
)

func Connect() {
	startViper()
	startGorm()
	startCasbin()
}

func startCasbin() {
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	// getting db for casbin
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	casbinDb, err := sql.Open("postgres", connectionString)
	if err != nil {
		llog.Error(err)
	}

	tableName := "casbin"
	adapter, err := casbinpgadapter.NewAdapter(casbinDb, tableName)
	if err != nil {
		llog.Error(err)
	}

	enforcer, err := casbin.NewEnforcer(ConfigPath+"/model.conf", adapter)
	if err != nil {
		llog.Error(err)
	}

	Casbin = enforcer
}

func startGorm() {
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	// getting db for gorm
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		llog.Error(err)
	}

	DB = db
}

func startViper() {
	llog.Info(fmt.Sprintf("Setting conf to config%s", RunMode))
	viper.SetConfigName(fmt.Sprintf("config%s", RunMode))
	viper.SetConfigType("toml")
	if ConfigPath == "" {
		ConfigPath = "./configs/"
	}
	viper.AddConfigPath(ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		llog.Notice(err)
	}

	viper.SetDefault("server.cert", "")
	viper.SetDefault("server.key", "")
	viper.SetDefault("options.debugMessages", false)

	Dm = viper.GetBool("options.debugMessages")

}
