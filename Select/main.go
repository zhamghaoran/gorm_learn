package main

import (
	"fmt"
	"gorm_learn"
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

// InsertInto 添加元素
func InsertInto() int64 {
	db := Utils.CreateDB()
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
	db := Utils.CreateDB()
	var student STUDENT
	db.First(&student)
	println(Utils.ToJsonString(student))

}

// SelectOne SELECT * FROM users LIMIT 1;
func SelectOne() {
	db := Utils.CreateDB()
	var student STUDENT
	db.Take(&student)
	fmt.Println(Utils.ToJsonString(student))

}

// SelectAll SELECT * FROM users
// 声明一个数组，调用find方法就可以查询到所有信息
func SelectAll() {
	db := Utils.CreateDB()
	var students []STUDENT
	db.Find(&students)
	fmt.Println(Utils.ToJsonString(students))
}

func main() {
	//InsertInto()
	//SelectLimit()
	//SelectOne()
	SelectAll()
}
