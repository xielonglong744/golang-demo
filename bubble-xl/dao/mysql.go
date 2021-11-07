package dao

import (
	"bubble-xl/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init(cfg *settings.GinConf) (err error){
	dsn := fmt.Sprintf("%s:%v@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MysqlConf.Username,cfg.MysqlConf.Password,cfg.MysqlConf.Hosts,cfg.MysqlConf.Port,cfg.MysqlConf.DB)
	DB,err = gorm.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("gorm open db failed %v\n",err)
		return err
	}
	fmt.Println("db init sucess")
	return nil
}

func Close() {
	DB.Close()
}
