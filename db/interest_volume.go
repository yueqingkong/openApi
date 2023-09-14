package db

import "time"

// 获取交割永续的持仓量和交易量
type InterestVolume struct {
	Id        int64
	Plat      string    `xorm:"plat"` // 平台名称
	Ccy       string    `xorm:"ccy" comment:"coin BTC/ETH"`
	Period    string    `xorm:"period" comment:"类型 swap"`
	Times     string    `xorm:"times" comment:"周期 5m/1H/1D"`
	Timestamp int64     `xorm:"time_stamp" comment:"数据产生时间"`
	Hold      float32   `xorm:"hold" comment:"持仓量(USD)"`
	Vol       float32   `xorm:"vol" comment:"交易量(USD)"`
	StartTime time.Time `xorm:"start_time"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
