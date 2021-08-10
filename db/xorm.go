package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var engine *xorm.Engine

// 连接数据库
func ConnectSQL(name,user,host,port,password string) {
	var err error

	// mysql配置
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, name)
	log.Print(sourceName)

	engine, err = xorm.NewEngine("mysql", sourceName)
	if err != nil {
		log.Fatal("[MySql] 连接失败,", err)
	}

	engine.ShowSQL(true)
	err = engine.Sync2(new(Coin), new(Account), new(Record))
	if err != nil {
		log.Fatal("[MySql] 同步表失败", err)
	}
}

func Engine() *xorm.Engine {
	return engine
}
