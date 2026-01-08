package model

import (
	"fmt"
	"github.com/v03413/bepusdt/app/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() error {
	dsn := conf.GetPostgresDSN()
	if dsn == "" {
		return fmt.Errorf("PostgreSQL DSN 不能为空")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("PostgreSQL 数据库初始化失败：%w", err)
	}

	if err = AutoMigrate(); err != nil {
		return fmt.Errorf("数据库结构迁移失败：%w", err)
	}

	addStartWalletAddress()

	return nil
}

func AutoMigrate() error {

	return DB.AutoMigrate(&WalletAddress{}, &TradeOrders{}, &NotifyRecord{}, &Config{}, &Webhook{})
}
