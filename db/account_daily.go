package db

import (
	"github.com/yueqingkong/openApi/conset"
	"log"
	"time"
	"xorm.io/builder"
)

type AccountDaily struct {
	Id         int64
	Plat       string    `xorm:"varchar(255)"`                            // 平台名称
	Symbol     string    `xorm:"varchar(255)"`                            // Token
	Total      float32   `xorm:"float"`                                   // 总值
	CreateTime time.Time `json:"create_time" xorm:"create_time DateTime"` // 时间
}

func (self *AccountDaily) Inserts(pt conset.PLAT, symbol conset.SYMBOL) error {
	self.Plat = dPlat(pt)
	self.Symbol = dSymbol(symbol)

	_, err := Engine().InsertOne(self)
	return err
}

// 当月账户余额记录
func (self *AccountDaily) MonthAccountBalances(pt conset.PLAT, symbol conset.SYMBOL) ([]*AccountDaily, error) {
	self.Plat = dPlat(pt)
	self.Symbol = dSymbol(symbol)

	balances := make([]*AccountDaily, 0)

	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), 0, 0, 0, 0, 0, currentTime.Location())
	sql, args, _ := builder.ToSQL(builder.Gte{"create_time": startTime})
	if err := Engine().Where(sql, args...).Find(&balances, self); err != nil {
		log.Printf("MonthAccountBalances err: %v", err)
		return nil, err
	}
	return balances, nil
}
