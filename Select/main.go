package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type STUDENT struct {
	SNO    string `gorm:"column:SNO"`
	SNAME  string `gorm:"column:SNAME"`
	SSEX   string `gorm:"column:SSEX"`
	SMAJOR string `gorm:"column:SMAJOR"`
	SDEPT  string `gorm:"column:SDEPT"`
	SAGE   int    `gorm:"column:SAGE"`
	TEL    int    `gorm:"column:TEL"`
	EAMIL  string `gorm:"column:EMAIL"`
}

func (STUDENT) TableName() string {
	return "STUDENT"
}
func ToJsonString(T any) string {
	jsonBytes, _ := json.Marshal(T)
	return string(jsonBytes)
}

// CreateDB 获取DB对象
func CreateDB() *gorm.DB {
	dsn := "username:password@@tcp(url)/STUDENTDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// InsertInto 添加元素
func InsertInto() int64 {
	db := CreateDB()
	student := STUDENT{
		SNO:    "123",
		SNAME:  "456",
		SSEX:   "男",
		SMAJOR: "101112",
		SDEPT:  "131415",
		SAGE:   9,
		TEL:    110,
		EAMIL:  "1123@qq.com",
	}
	result := db.Create(&student)
	return result.RowsAffected
}

// SelectLimit SelectOne SELECT * FROM users ORDER BY id LIMIT 1;
func SelectLimit() {
	db := CreateDB()
	var student STUDENT
	db.First(&student)
	println(ToJsonString(student))

}

// SelectOne SELECT * FROM users LIMIT 1;
func SelectOne() {
	db := CreateDB()
	var student STUDENT
	db.Take(&student)
	fmt.Println(ToJsonString(student))

}

// SelectAll SELECT * FROM users
// 声明一个数组，调用find方法就可以查询到所有信息
func SelectAll() {
	db := CreateDB()
	var students []STUDENT
	db.Find(&students)
	fmt.Println(ToJsonString(students))
}
func main() {
	//InsertInto()
	//SelectLimit()
	//SelectOne()
	//SelectAll()
}
