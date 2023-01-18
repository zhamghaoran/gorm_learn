package utils

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ToJsonString(T any) string {
	jsonBytes, _ := json.Marshal(T)
	return string(jsonBytes)
}

// CreateDB 获取DB对象
func CreateDB() *gorm.DB {
	dsn := "username:password@tcp(url)/STUDENTDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
