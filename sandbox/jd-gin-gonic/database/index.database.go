package database

import (
	"fmt"
	"log"

	"github.com/Yuseonr/gin-labs/sandbox/jd-gin-gonic/config/db_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	var err error

	switch db_config.DB_DRIVE {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD,db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME,
		)

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	case "pgsql":
		return fmt.Errorf("postgres not implemented yet")

	default:
		return fmt.Errorf("unsupported database driver: %s", db_config.DB_DRIVE)
	}

	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}

	log.Println("Connected to database")
	return nil
}
