package psql

import (
	"log"
	"login/model"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var pdb *gorm.DB
var store model.Store
var storeOnce sync.Once

type Store struct {
	db *gorm.DB
}

func InitPsqlDB() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	dsn := "host=127.0.0.1 user=postgres password=123456@ dbname=account port=5432 sslmode=disable "
	PsqlHelper, psqlErr := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if psqlErr != nil {
		return psqlErr
	}

	sqlDB, _ := PsqlHelper.DB()
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	pdb.Logger.LogMode(logger.Info)
	createTime := time.Now()
	user := model.User{
		UserName:    "lx",
		PassWord:    "123456",
		Email:       "163@163.com",
		Phone:       13000000000,
		CreatedTime: createTime,
	}
	pdb.Create(&user)
	pdb.AutoMigrate(&model.User{})
	return nil
}

func SharedStore() model.Store {
	storeOnce.Do(func() {
		err := InitPsqlDB()
		if err != nil {
			panic(err)
		}
		store = NewStore(pdb)
	})
	return store
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) BeginTx() (model.Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) CommitTx() error {
	return s.db.Commit().Error
}
