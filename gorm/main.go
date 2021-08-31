package main

import (
	"component-study1/gorm/global"
	"component-study1/gorm/model"
	"fmt"
	"log"
)

func init(){
	err := setupDBEngine()
	if err != nil{
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	fmt.Println(global.DBEngine)
}

func setupDBEngine() error{
	var err error
	global.DBEngine, err = model.NewDBEngine()
	if err != nil{
		return err
	}
	return nil
}