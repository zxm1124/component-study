package main

import (
	"component-study1/viper/global"
	"component-study1/viper/pkg/setting"
	"fmt"
	"log"
)

func init(){
	err := setupSetting()
	if err != nil{
		log.Fatalf("setupSetting err: %v",err)
	}
}

func main() {
	fmt.Println(global.APPSetting)
}

// 读取配置文件信息
func setupSetting() error{
	setting, err := setting.NewSetting()
	if err != nil{
		return err
	}
	// 读取APP配置文件信息
	err = setting.ReadSection("APP", &global.APPSetting)
	if err != nil{
		log.Fatalf("setupSetting err: %v",err)
		return err
	}
	return nil
}
