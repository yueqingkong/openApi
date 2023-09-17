package db

import "time"

// 精英多空平均持仓比例
type TopAverageIndex struct {
	Id        int64
	Plat      string    `xorm:"plat"` // 平台名称
	Ccy       string    `xorm:"ccy" comment:"coin BTC/ETH"`
	Period    string    `xorm:"period" comment:"类型 CONTRACTS"`
	Times     string    `xorm:"times" comment:"周期 5m/1H/1D"`
	Timestamp int64     `xorm:"time_stamp" comment:"数据产生时间"`
	Long      float32   `xorm:"long" comment:"做多持仓比例"`
	Short     float32   `xorm:"short" comment:"做空持仓比例"`
	StartTime time.Time `xorm:"start_time"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
