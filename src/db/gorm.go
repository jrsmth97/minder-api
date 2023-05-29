package db

import (
	"fmt"
	"minder/src/server/model"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGormDB() (*gorm.DB, error) {
	var (
		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
		DB_USER = os.Getenv("DB_USER")
		DB_PASS = os.Getenv("DB_PASS")
		DB_NAME = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbs, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = dbs.Ping()
	if err != nil {
		return nil, err
	}

	err = db.Debug().AutoMigrate(
		model.User{},
		model.UserInterest{},
		model.UserMembership{},
		model.UserPhoto{},
		model.UserSwipe{},
		model.Location{},
		model.Interest{},
		model.Privilege{},
		model.Purchase{},
		model.Membership{},
		model.MembershipPrivilege{},
	)

	return db, nil
}
