package db

import "time"

// 主动买入和卖出的交易量
type TakeVolume struct {
	Id        int64
	Plat      string    `xorm:"plat"` // 平台名称
	Ccy       string    `xorm:"ccy" comment:"coin BTC/ETH"`
	Period    string    `xorm:"period" comment:"类型 SPOT/CONTRACTS"`
	Times     string    `xorm:"times" comment:"周期 5m/1H/1D"`
	Timestamp int64     `xorm:"time_stamp" comment:"数据产生时间"`
	Buy       float32   `xorm:"buy" comment:"买入量"`
	Sell      float32   `xorm:"sell" comment:"卖出量"`
	StartTime time.Time `xorm:"start_time"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
