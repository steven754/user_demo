package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"test/setting"
)

var (
	DB *gorm.DB
)

func ConnectMysql(cfg *setting.MySQLConfig) (err error) {
	//dsn1 := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Database)
	//fmt.Printf(dsn1)
	dsn := "root:123456@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql连接失败，请检查", err)
		log.Fatal(err)
		return
	}
	fmt.Println("mysql连接成功")
	return
}
