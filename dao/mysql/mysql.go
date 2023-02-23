package mysql

import (
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

//const dbName = "pwh:123456@tcp(47.106.206.78:3306)/system_admin?charset=utf8mb4&parseTime=true"

/*func InitializeDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
	}
}*/

func InitializeDB(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB.Use(dbresolver.Register(dbresolver.Config{}).SetMaxOpenConns(cfg.MaxOpenConns).SetMaxIdleConns(cfg.MaxIdleConns))
	return
}
