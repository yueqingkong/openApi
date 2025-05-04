package db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

// 连接数据库
func ConnectSQL(name, user, host, port, password string, showSQL bool) {
	var err error

	// mysql配置
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, name)
	log.Print(sourceName)

	engine, err = xorm.NewEngine("mysql", sourceName)
	if err != nil {
		log.Fatal("[MySql] 连接失败,", err)
	}

	engine.ShowSQL(showSQL)
	engine.DatabaseTZ = time.Local
	engine.TZLocation = time.Local
	err = engine.Sync2(
		&Coin{},
		&Account{},
		&Record{},
		&AccountDay{},
		&Indexs{},
		&Instrument{},
		&OLog{},
		&Address{},
		&AddressTransaction{},
		&FundingRate{},
		&TakeVolume{},
		&LoanRatio{},
		&AccountRate{},
		&InterestVolume{},
		&TopSentimentIndex{},
		&TopAverageIndex{})
	if err != nil {
		log.Fatal("[MySql] 同步表失败", err)
	}
}

func Engine() *xorm.Engine {
	return engine
}
