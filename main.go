package main

import (
	"fmt"
	"test/domain"
	"test/models"
	"test/mysql"
	"test/setting"
)

func main() {
	//命令行启动并指定配置文件
	//if len(os.Args) < 2 {
	//	fmt.Println("Usage：./conf conf/config.ini")
	//	return
	//}

	// 加载配置文件
	if err := setting.Init("conf/config.ini"); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库

	// 连接数据库
	err := mysql.ConnectMysql(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("connect mysql failed, err:%v\n", err)
		return
	}
	//defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	err = mysql.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("数据库自动迁移错误")
		return
	}
	mysql.DB.Migrator().CurrentDatabase()

	// 注册路由
	r := domain.SetupRouter()
	fmt.Println(setting.Conf.Port)
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
