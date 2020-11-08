package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"zx/untils"
	"zx/untils_model"
	"zx/untils_output"
)

func Init() *gorm.DB {
	err := untils.GetConfigIni("./config/database.ini")
	if err != nil {
		panic(err)
	}
	//读取配置文件
	var input untils_model.Input
	output, _ := untils_output.Function(input)
	//fmt.Println(output.Database.Local)
	db, err2 := gorm.Open(output.Database.Types, output.Database.Local)
	//if err1 != nil {
	//	fmt.Println("读取配置文件出错了:%v",err1)
	//}
	if err2 != nil {
		fmt.Println("数据库连接错误:%v", err2)
	}

	return db

}
