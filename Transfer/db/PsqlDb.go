package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"transferasset/model"

	"github.com/kjk/betterguid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PsqlHelper *gorm.DB

var psqlErr error

func InitPsqlDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	dsn := "host=127.0.0.1 user=postgres password=123456@ dbname=account port=5432 sslmode=disable "
	PsqlHelper, psqlErr := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if psqlErr != nil {
		fmt.Println(psqlErr)
	}
	sqlDB, _ := PsqlHelper.DB()
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	PsqlHelper.Logger.LogMode(logger.Info)
	createAt := time.Now()
	asset := model.Asset{
		Dto: model.Dto{
			Id:         betterguid.New(),
			CreateTime: createAt.Unix(),
			CreatorId:  "admin",
			Deleted:    false,
		},
		CustomerId:        "test001",
		BusinessId:        "1",
		CoinId:            "1",
		AvaliableQuantity: 100, //可用
		FrozenQuantity:    0,   //占用
		Include:           0,   //划入
		Drawout:           0,   //划出
	}
	PsqlHelper.Create(&asset)
	PsqlHelper.AutoMigrate(&model.Asset{}, &model.AssetWasteBook{}, &model.AssetTransferRecord{})
}
