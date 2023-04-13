package db

import (
	"errors"
	"github.com/yueqingkong/openApi/conset"
	"time"
	"xorm.io/builder"
)

// Record 交易记录
type Record struct {
	Id            int64
	Name          string    `xorm:"name index(name-symbol-period) index(n-s-s) index(n-s-o-p)"`   // 策略
	Plat          string    `xorm:"plat"`                                                         // 平台名称
	Symbol        string    `xorm:"symbol index(name-symbol-period) index(n-s-s) index(n-s-o-p)"` // Token
	Period        string    `xorm:"period index(name-symbol-period)"`                             // 周期 spot|week|quarter
	Operation     int32     `xorm:"int index(n-s-o-p)"`                                           // 1: 开多 2: 开空 3: 平仓
	Position      int32     `xorm:"int index(n-s-o-p)"`                                           // 加仓层数
	Price         float32   `xorm:"float"`                                                        // 当前价格
	AvgPrice      float32   `xorm:"float"`                                                        // 均价
	Used          float32   `xorm:"float"`                                                        // 已开仓Token
	Size          float32   `xorm:"float"`                                                        // 开仓张数
	Total         float32   `xorm:"float"`                                                        // 当前账户总值
	LossPrice     float32   `xorm:"float"`                                                        // 止损价
	EstimatedLoss float32   `xorm:"float"`                                                        // 预计亏损
	Detail        string    `xorm:"detail text"`                                                  // 描述 usd->token | ust<-token
	Profit        float32   `xorm:"float"`                                                        // 收益
	ProfitRate    float32   `xorm:"float"`                                                        // 收益率(百分比 %)
	TotalRate     float32   `xorm:"float"`                                                        // 总收益率(百分比 %)
	Status        int32     `xorm:"status index(n-s-s)"`                                          // 状态
	CreateTime    time.Time `json:"create_time" xorm:"create_time index"`                // 时间
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
}

func Operation(op int) string {
	var s string
	if op == conset.BUY_HIGH {
		s = "开多"
	} else if op == conset.BUY_LOW {
		s = "开空"
	} else if op == conset.SELL_HIGH {
		s = "平多"
	} else if op == conset.SELL_LOW {
		s = "平空"
	}
	return s
}

func (self *Record) Insert(pt conset.PLAT, base conset.CCY, quote conset.CCY) error {
	self.Plat = Plat(pt)
	self.Symbol = Symbol(base, quote)

	_, err := Engine().InsertOne(self)
	return err
}

func (self *Record) Last(pt conset.PLAT, base conset.CCY, quote conset.CCY) error {
	self.Plat = Plat(pt)
	self.Symbol = Symbol(base, quote)

	if b, err := Engine().Desc("create_time").Get(self); err != nil || !b {
		return errors.New("get")
	}

	return nil
}

func (self *Record) Clear() error {
	sql, args, _ := builder.ToSQL(builder.Gte{"id": 0})
	_, err := Engine().Where(sql, args...).Delete(self)
	return err
}
