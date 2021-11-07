package settings

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var Conf = new(GinConf)

type GinConf struct {
	MysqlConf `ini:"mysql"`
	AppConf `ini:"gininfo"`
}


type MysqlConf struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Port int `ini:"port"`
	Hosts string `ini:"hosts"`
	DB string `ini:"db"`
}

type AppConf struct {
	Port int `ini:"port"`
	Release bool `ini:"release"`
}

func LoadConf(filename string) error{
	err := ini.MapTo(Conf,filename)
	if err != nil {
		fmt.Printf("load file failed err %v\n", err)
		return err
	}
	return nil
}