package database

import (
	"fmt"
	"github.com/hoangnguyen-1312/faucet/domain"
	"github.com/hoangnguyen-1312/faucet/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

func InitDatabase(conn config.DatabaseConnection) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conn.User, conn.Password, conn.Host, conn.Port, conn.Dbname))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect successfully")
	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&domain.TransactionLog{})
}