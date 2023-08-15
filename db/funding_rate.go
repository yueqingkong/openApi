package db

import "time"

// 资金费率
type FundingRate struct {
	Id        int64
	Symbol    string    `xorm:"symbol"`
	Rate      string    `xorm:"rate"`
	StartTime string    `xorm:"start_time"`
	NextRate  string    `xorm:"next_rate"`
	NextTime  string    `xorm:"next_time"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
