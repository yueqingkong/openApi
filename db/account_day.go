package db

import (
	"github.com/yueqingkong/openApi/conset"
	"time"
)

type AccountDay struct {
	Id        int64
	DayTime   string    `xorm:"day_time" comment:"日期"`
	Name      string    `xorm:"varchar(255)"` // 名称
	Plat      string    `xorm:"varchar(255)"` // 平台名称
	Symbol    string    `xorm:"varchar(255)"` // Token
	Used      float32   `xorm:"float"`        // 总值
	Total     float32   `xorm:"float"`        // 总值
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (self *AccountDay) Inserts(pt conset.PLAT, base conset.CCY, quote conset.CCY) error {
	self.Plat = Plat(pt)
	self.Symbol = Symbol(base, quote)

	_, err := Engine().InsertOne(self)
	return err
}
