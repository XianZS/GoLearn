package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DBconnect *gorm.DB

var err error

func DD() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "XianZS:3135QWE@tcp(127.0.0.1:3306)/golandlearning?charset=utf8mb4&parseTime=True&loc=Local"
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		fmt.Println(DBconnect)
	}
}

// func init() {
// 	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
