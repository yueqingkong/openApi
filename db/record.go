package db

import (
	"errors"
	"github.com/yueqingkong/openApi/conset"
	"log"
	"time"
	"xorm.io/builder"
)

// Record 交易记录
type Record struct {
	Id            int64
	Plat          string    `xorm:"varchar(255) plat index(plat-symbol-period)"`   // 平台名称
	Symbol        string    `xorm:"varchar(255) symbol index(plat-symbol-period)"` // Token
	Period        string    `xorm:"varchar(255) period index(plat-symbol-period)"` // 周期 spot|week|quarter
	Operation     int32     `xorm:"int"`                                           // 1: 开多 2: 开空 3: 平仓
	Position      int32     `xorm:"int"`                                           // 加仓层数
	Price         float32   `xorm:"float"`                                         // 当前价格
	AvgPrice      float32   `xorm:"float"`                                         // 均价
	Used          float32   `xorm:"float"`                                         // 已开仓Token
	Size          float32   `xorm:"float"`                                         // 开仓张数
	Total         float32   `xorm:"float"`                                         // 当前账户总值
	LossPrice     float32   `xorm:"float"`                                         // 止损价
	EstimatedLoss float32   `xorm:"float"`                                         // 预计亏损
	Detail        string    `xorm:"detail text"`                                   // 描述 usd->token | ust<-token
	Profit        float32   `xorm:"float"`                                         // 收益
	ProfitRate    float32   `xorm:"float"`                                         // 收益率
	TotalRate     float32   `xorm:"float"`                                         // 总收益率
	CreateTime    time.Time `json:"create_time" xorm:"create_time DateTime index"` // 时间
}

func (self *Record) Insert(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD) error {
	self.Plat = dPlat(pt)
	self.Symbol = dSymbol(symbol)
	self.Period = dPeriod(period)

	_, err := Engine().InsertOne(self)
	return err
}

func (self *Record) Last(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD) error {
	self.Plat = dPlat(pt)
	self.Symbol = dSymbol(symbol)
	self.Period = dPeriod(period)

	if b, err := Engine().Desc("create_time").Get(self); err != nil || !b {
		return errors.New("get")
	}

	return nil
}

// 最近首次开仓
func (self *Record) LastFirstOpenRecord(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD) error {
	self.Plat = dPlat(pt)
	self.Symbol = dSymbol(symbol)
	self.Period = dPeriod(period)

	sql, args, _ := builder.ToSQL(builder.In("operation", conset.BUY_HIGH, conset.BUY_LOW).
		And(builder.Eq{"position": 1}))
	if b, err := Engine().Where(sql, args...).Desc("create_time").Get(self); err != nil || !b {
		return errors.New("get")
	}

	return nil
}

// 最近未平仓
func (self *Record) RecentOpenRecords(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD) ([]*Record, error) {
	if err := self.LastFirstOpenRecord(pt, symbol, period); err != nil {
		return nil, err
	}
	records := make([]*Record, 0)

	sql, args, _ := builder.ToSQL(builder.Gte{"id": self.Id})
	if err := Engine().Where(sql, args...).Asc("create_time").Find(&records, &Record{}); err != nil {
		return nil, err
	}

	return records, nil
}

// 昨日平仓
func (self *Record) YesTodayCloseRecords(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD) ([]*Record, error) {
	self.Plat = dPlat(pt)
	self.Symbol = dSymbol(symbol)
	self.Period = dPeriod(period)

	records := make([]*Record, 0)

	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-1, 0, 0, 0, 0, currentTime.Location())
	sql, args, _ := builder.ToSQL(builder.Gte{"create_time": startTime})
	if err := Engine().Where(sql, args...).Find(&records, self); err != nil {
		log.Printf("history closeRecords err: %v", err)
		return nil, err
	}
	return records, nil
}

func (self *Record) Clear() error {
	sql, args, _ := builder.ToSQL(builder.Gte{"id": 0})
	_, err := Engine().Where(sql, args...).Delete(self)
	return err
}
