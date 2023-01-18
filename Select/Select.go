package Select

import (
	"fmt"
	"gorm_learn/utils"
)

// 获取数据库连接对象
var db = utils.CreateDB()

type STUDENT struct {
	SNO    string `gorm:"column:SNO;primaryKey"`
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
	var student STUDENT
	db.First(&student)
	println(utils.ToJsonString(student))

}

// SelectOne SELECT * FROM users LIMIT 1;
func SelectOne() {
	var student STUDENT
	db.Take(&student)
	fmt.Println(utils.ToJsonString(student))

}

// SelectAll SELECT * FROM users
// 声明一个数组，调用find方法就可以查询到所有信息
func SelectAll() {
	var students []STUDENT
	db.Find(&students)
	fmt.Println(utils.ToJsonString(students))
}

// StringConditionSelect Select * from STUDENT where SNO = '123' and SNAME = '456'
func StringConditionSelect() {
	var students []STUDENT
	db.Where("SNO = ? AND SNAME = ?", "123", "456").Find(&students)
	fmt.Println(utils.ToJsonString(students))
}

// SelectRow Select SNO,SNAME from STUDENT
func SelectRow() {
	var students []STUDENT
	db.Select("SNO", "SNAME").Find(&students)
	fmt.Println(utils.ToJsonString(students))
}
