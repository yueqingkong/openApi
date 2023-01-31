package db

import (
	"errors"
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/util"
	"time"
)

type Indexs struct {
	Id        int64
	Plat      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 平台名称
	Name      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 指标名称
	Symbol    string    `xorm:"varchar(255) unique(p-n-s-d)"` // 币种
	Date      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 格式化时间
	P1        float32   `xorm:"p_1"`
	P2        float32   `xorm:"p_2"`
	P3        float32   `xorm:"p_3"`
	P4        float32   `xorm:"p_4"`
	P5        float32   `xorm:"p_5"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (self *Indexs) Get(pt conset.PLAT, name string, bs conset.CCY, quote conset.CCY, start time.Time) (*Indexs, error) {
	indexs := &Indexs{Plat: Plat(pt), Name: name, Symbol: Symbol(bs, quote), Date: util.TimeFormatDay(start)}
	if b, err := Engine().Get(indexs); err != nil {
		return nil, err
	} else if !b {
		return nil, errors.New("not exist")
	}

	return indexs, nil
}

func (self *Indexs) Create(pt conset.PLAT, name string, bs conset.CCY, quote conset.CCY, start time.Time, short, long, low, high, atr float32) error {
	record := &Indexs{Plat: Plat(pt), Name: name, Symbol: Symbol(bs, quote), Date: util.TimeFormatDay(start),
		P1: short, P2: long, P3: low, P4: high, P5: atr}

	_, err := Engine().InsertOne(record)
	return err
}
