package db

import (
	"errors"
	"fmt"
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/util"
	"strings"
	"time"
)

// Instrument 交易产品基础信息
type Instrument struct {
	Id        int64
	Plat      string    `xorm:"plat"`      // 平台名称
	Period    string    `xorm:"period"`    // spot
	Symbol    string    `xorm:"symbol"`    // btc_usdt
	CtVal     float32   `xorm:"ct_val"`    // 合约面值，仅适用于交割/永续/期权
	State     string    `xorm:"state"`     // 产品状态 live：交易中 suspend：暂停中 preopen：预上线 test：测试中（测试产品，不可交易）
	ListTime  time.Time `xorm:"list_time"` // 上线日期
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (self *Instrument) Get(ins *Instrument) error {
	if b, err := Engine().Get(ins); err != nil {
		return err
	} else if !b {
		return errors.New("not exist")
	}
	return nil
}

func (self *Instrument) Create(pt conset.PLAT, period conset.PERIOD, instId, uly, instFamily, settleCcy, ctVal, state string) error {
	ccies := strings.Split(instId, "-")
	_, err := Engine().InsertOne(&Instrument{
		Plat:   Plat(pt),
		Period: Period(period),
		Symbol: fmt.Sprintf("%s_%s", strings.ToLower(ccies[0]), strings.ToLower(ccies[1])),
		CtVal:  util.Float32(ctVal),
		State:  state,
	})
	return err
}

func (self *Instrument) Update(id int64, ctVal, state string) error {
	instrument := &Instrument{CtVal: util.Float32(ctVal), State: state}
	if _, err := Engine().Update(instrument, &Instrument{Id: id}); err != nil {
		return err
	}
	return nil
}

func (self *Instrument) CreateOrUpdate(pt conset.PLAT, period conset.PERIOD, instId, uly, instFamily, settleCcy, ctVal, state string) error {
	ccies := strings.Split(instId, "-")
	ins := &Instrument{Plat: Plat(pt), Period: Period(period), Symbol: fmt.Sprintf("%s_%s", strings.ToLower(ccies[0]), strings.ToLower(ccies[1]))}
	if err := self.Get(ins); err != nil {
		return self.Create(pt, period, instId, uly, instFamily, settleCcy, ctVal, state)
	} else {
		val := util.Float32(ctVal)
		if val != ins.CtVal || state != ins.State {
			return self.Update(ins.Id, ctVal, state)
		}
	}
	return nil
}
