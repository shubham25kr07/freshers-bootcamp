package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DataBaseConfig struct {
	User     string
	Password string
	DbName   string
}

func BuildConfig() *DataBaseConfig {
	dbConfig := DataBaseConfig{
		User:     "root",
		Password: "***",
		DbName:   "day4and5",
	}
	return &dbConfig
}

func DBurl(db *DataBaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.DbName,
	)
}
