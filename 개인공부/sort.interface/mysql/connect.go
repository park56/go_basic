package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StockInfo struct {
	Name  string `gorm:"column:item_name"`
	Price string `gorm:"column:current_price"`  // 현재가
	PDE   string `gorm:"column:then_yesterday"` // 전일비
	TV    string `gorm:"column:trading_volume"` // 거래량
	MC    string `gorm:"column:market_cap"`     // 시가총액
}

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := "root:1234@tcp(127.0.0.1:3306)/hack?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	DB = db

	return db
}

func GetStockInfo() []*StockInfo {
	SI := []*StockInfo{}
	if err := DB.Limit(10).Find(&SI).Error; err != nil {
		log.Println(err)
	}
	return SI
}
