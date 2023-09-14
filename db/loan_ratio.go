package db

import "time"

// 获取借入计价货币与借入交易货币的累计数额比值
type LoanRatio struct {
	Id        int64
	Plat      string    `xorm:"plat"` // 平台名称
	Ccy       string    `xorm:"ccy" comment:"coin BTC/ETH"`
	Times     string    `xorm:"times" comment:"周期 5m/1H/1D"`
	Timestamp int64     `xorm:"time_stamp" comment:"数据产生时间"`
	Ratio     float32   `xorm:"ratio" comment:"多空比值"`
	StartTime time.Time `xorm:"start_time"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
