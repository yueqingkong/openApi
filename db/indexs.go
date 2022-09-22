package db

import "time"

type Indexs struct {
	Id        int64
	Plat      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 平台名称
	Name      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 指标名称
	Symbol    string    `xorm:"varchar(255) unique(p-n-s-d)"` // 币种
	Date      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 格式化时间
	P1        float32   `xorm:"p_1 "`
	P2        float32   `xorm:"p_2"`
	P3        float32   `xorm:"p_3"`
	P4        float32   `xorm:"p_4"`
	P5        float32   `xorm:"p_5"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
