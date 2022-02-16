package db

import (
	"errors"
	"github.com/yueqingkong/openApi/conset"
	"xorm.io/builder"
)

type Account struct {
	Id     int64
	Plat   string  `xorm:"varchar(255)"` // 平台名称
	Symbol string  `xorm:"varchar(255)"` // Token
	Total  float32 `xorm:"float"`        // 总值
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
	if _, err := Engine().Update(self, &Account{Id: self.Id}); err != nil {
		return err
	}

	return nil
}

func (self *Account) Clear() error {
	sql, args, _ := builder.ToSQL(builder.Gte{"id": 0})
	_, err := Engine().Where(sql, args...).Delete(self)
	return err
}
