package db

import (
	"github.com/yueqingkong/openApi/conset"
	"time"
)

type AccountDay struct {
	Id         int64
	Name       string    `xorm:"varchar(255)"`                            // 名称
	Plat       string    `xorm:"varchar(255)"`                            // 平台名称
	Symbol     string    `xorm:"varchar(255)"`                            // Token
	Used       float32   `xorm:"float"`                                   // 总值
	Total      float32   `xorm:"float"`                                   // 总值
	CreateTime time.Time `json:"create_time" xorm:"create_time DateTime"` // 时间
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

func (self *AccountDay) Inserts(pt conset.PLAT, symbol conset.SYMBOL) error {
	self.Plat = dPlat(pt)
	self.Symbol = SymbolToString(symbol)

	_, err := Engine().InsertOne(self)
	return err
}
