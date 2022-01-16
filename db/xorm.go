package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

var engine *xorm.Engine

// 连接数据库
func ConnectSQL(name,user,host,port,password string) {
	var err error

	// mysql配置
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, port, name)
	log.Print(sourceName)

	engine, err = xorm.NewEngine("mysql", sourceName)
	if err != nil {
		log.Fatal("[MySql] 连接失败,", err)
	}

	engine.ShowSQL(true)
	// 本地时区
	engine.DatabaseTZ = time.Local // 必须
	engine.TZLocation = time.Local // 必须
	err = engine.Sync2(new(Coin), new(Account), new(Record),new(AccountDay))
	if err != nil {
		log.Fatal("[MySql] 同步表失败", err)
	}
}

func Engine() *xorm.Engine {
	return engine
}
