package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"transferasset/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBHelper *gorm.DB
var err error

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)

	dsn := "root:1230@/transfer?charset=utf8mb4&parseTime=True&loc=Local" //给个错误看看启动效果

	DBHelper, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := DBHelper.DB()
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	DBHelper.Logger.LogMode(logger.Info)

	DBHelper.AutoMigrate(&model.Asset{}, &model.AssetWasteBook{}, &model.AssetTransferRecord{})
}
