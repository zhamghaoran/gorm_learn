package Update

import (
	"fmt"
	"gorm_learn/Select"
	Utils "gorm_learn/utils"
)

var db = Utils.CreateDB()

// Update 根据主键来构建where 条件，所以一定要指定主键，不然将会无法更新
func Update() {
	var student Select.STUDENT
	db.First(&student)
	student.EAMIL = "11112@qq.com"
	tx := db.Save(&student)
	fmt.Println(tx.RowsAffected)
}

// UpdateByCondition Update STUDENT set SNAME = 'qaq' where SNAME = '李素素'
func UpdateByCondition() {
	db.Model(Select.STUDENT{}).Where("SNAME = ?", "李素素").Update("SNAME", "qaq")
}

// Delete  DELETE FROM `STUDENT` WHERE SNAME = '456'
func Delete() {
	var student Select.STUDENT
	db.First(&student)
	db.Where("SNAME = ?", "456").Delete(&Select.STUDENT{})

}
