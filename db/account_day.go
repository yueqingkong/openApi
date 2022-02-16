package db

import (
	"github.com/yueqingkong/openApi/conset"
	"time"
)

type AccountDay struct {
	Id         int64
	Plat       string    `xorm:"varchar(255)"`                            // 平台名称
	Symbol     string    `xorm:"varchar(255)"`                            // Token
	Total      float32   `xorm:"float"`                                   // 总值
	CreateTime time.Time `json:"create_time" xorm:"create_time DateTime"` // 时间
}

func (self *AccountDay) Inserts(pt conset.PLAT, symbol conset.SYMBOL) error {
	self.Plat = dPlat(pt)
	self.Symbol = SymbolToString(symbol)

	_, err := Engine().InsertOne(self)
	return err
}