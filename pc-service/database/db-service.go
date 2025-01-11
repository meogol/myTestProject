package database

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"meogol/pc-service/config"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func initConnection() (*gorm.DB, error) {
	var err error
	config := config.CurrentConfig
	dsn := "host=" + config.Database.DbHost + " user=" + config.Database.DbUser + " password=" + config.Database.DbPassword + " dbname=" + config.Database.DbName + " port=" + config.Database.DbPort + " sslmode=" + config.Database.SslMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database; Reason: %s")
	}

	return db, nil
}

func GetInstance() (*gorm.DB, error) {
	if db == nil {
		var err error
		db, err = initConnection()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func InitDatabase() error {
	var err error
	once.Do(func() {
		db, err = GetInstance()
		if err != nil {
			dbLogger.DPanicf("failed to connect to database; Reason: %s", err)
		}
	})
	return err
}
