package db

import (
	"github.com/yueqingkong/openApi/conset"
	"time"
)

type AddressTransaction struct {
	Id        int64
	Address   string    `xorm:"address"`
	BlockNum  int32     `xorm:"block_num"`
	From      string    `xorm:"from"`
	To        string    `xorm:"to"`
	Value     string    `xorm:"value"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (self *AddressTransaction) Last() error {
	if b, err := Engine().Desc("block_num").Get(self); err != nil {
		return err
	} else if !b {
		return conset.NOT_FOUND
	}
	return nil
}
