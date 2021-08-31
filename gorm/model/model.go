package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine() (*gorm.DB, error){
	url := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"root", "root",3306, "dorm", "utf8", "True")
	db, err := gorm.Open("mysql", url)
	if err != nil{
		return nil, err
	}
	// 判断是否未debug模式，如果是则启用
	db.LogMode(true)
	return db, err
}

