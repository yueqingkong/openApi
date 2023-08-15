package db

import (
	"time"
)

type Address struct {
	Id        int64
	Name      string    `xorm:"name"`
	Address   string    `xorm:"address"`
	Block     int32     `xorm:"block" comment:"跟踪到的block"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (self *Address) All() []*Address {
	items := make([]*Address, 0)
	Engine().Find(&items, &Address{})
	return items
}

func (self *Address) Update() error {
	if _, err := Engine().Update(self, &Address{Id: self.Id}); err != nil {
		return err
	}

	return nil
}
