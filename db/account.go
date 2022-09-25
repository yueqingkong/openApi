package db

import (
	"errors"
	"github.com/yueqingkong/openApi/conset"
	"time"
	"xorm.io/builder"
)

type Account struct {
	Id        int64
	Name      string    `xorm:"varchar(255) unique(n-p-s)"` // 名称
	Plat      string    `xorm:"varchar(255) unique(n-p-s)"` // 平台名称
	Symbol    string    `xorm:"varchar(255) unique(n-p-s)"` // Token
	Used      float32   `xorm:"float"`                      // 总值
	Total     float32   `xorm:"float"`                      // 总值
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (self *Account) Inserts(pt conset.PLAT, symbol conset.SYMBOL) error {
	self.Plat = dPlat(pt)
	self.Symbol = SymbolToString(symbol)

	_, err := Engine().InsertOne(self)
	return err
}

func (self *Account) Account(pt conset.PLAT, symbol conset.SYMBOL) error {
	self.Plat = dPlat(pt)
	self.Symbol = SymbolToString(symbol)

	if b, err := Engine().Get(self); err != nil || !b {
		return errors.New("get")
	}

	return nil
}

func (self *Account) Update() error {
	if _, err := Engine().Cols("symbol", "used", "total").Update(self, &Account{Id: self.Id}); err != nil {
		return err
	}

	return nil
}

func (self *Account) Clear() error {
	sql, args, _ := builder.ToSQL(builder.Gte{"id": 0})
	_, err := Engine().Where(sql, args...).Delete(self)
	return err
}
