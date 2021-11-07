package main

import (
	"bubble-xl/dao"
	"bubble-xl/modles"
	"bubble-xl/router"
	"bubble-xl/settings"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please input config path example ./conf/config.ini")
		return
	}
	filename := os.Args[1]
	// 1. 加载配置文件获取mysql连接信息
	err := settings.LoadConf(filename)
	if err != nil {
		fmt.Printf("settings load config failed %v\n", err)
		return
	}

	// 2. 初始化mysql
	err = dao.Init(settings.Conf)
	if err != nil {
		fmt.Printf("mysql init failed %v\n",err)
	}

	defer dao.Close()
	// 3. 数据库模型绑定
	dao.DB.AutoMigrate(&modles.Todo{})

	// 4. 注册路由
	r := router.Init(settings.Conf)
	err = r.Run(fmt.Sprintf(":%d",settings.Conf.AppConf.Port))
	if err != nil {
	}
}
