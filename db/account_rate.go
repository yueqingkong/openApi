package db

import "time"

// 获取交割永续净开多持仓用户数与净开空持仓用户数的比值
type AccountRate struct {
	Id        int64
	Plat      string    `xorm:"plat"` // 平台名称
	Ccy       string    `xorm:"ccy" comment:"coin BTC/ETH"`
	Times     string    `xorm:"times" comment:"周期 5m/1H/1D"`
	Period    string    `xorm:"period" comment:"类型 swap"`
	Timestamp int64     `xorm:"time_stamp" comment:"数据产生时间: 秒"`
	Ratio     float32   `xorm:"ratio" comment:"多空人数比"`
	StartTime time.Time `xorm:"start_time"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
