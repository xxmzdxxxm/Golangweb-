package loverwaitdb

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接池
var _db *gorm.DB
var dsn string = "root:root@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("connect error")
	}
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)

}

func GetDB() *gorm.DB {
	return _db
}
