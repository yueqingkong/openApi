package db

import "time"

type Index struct {
	Id         int64
	Plat       string    `xorm:"varchar(255)"`                   // 平台名称
	Name       string    `xorm:"varchar(255)"`                   //
	Index      float32   `xorm:"float"`                          // 指标
	CreateTime time.Time `json:"create_time" xorm:"create_time"` // 时间
}
