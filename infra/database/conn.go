package database

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"plataform_init/db"
)

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := "host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.Database + " port=" + config.Port + " sslmode=" + config.SSLMode
	fmt.Println(dsn)

	dbGorm, err := gorm.Open(driver.New(driver.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = dbGorm.AutoMigrate(&db.User{})
	if err != nil {
		return nil, err
	}

	return dbGorm, nil
}
