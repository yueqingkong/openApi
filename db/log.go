package db

import "time"

// 日志记录
type OLog struct {
	Id        int64
	Name      string    `xorm:"name" comment:"指标名称"`
	Task      string    `xorm:"task unique" comment:"任务标识"`
	Title     string    `xorm:"title" comment:"标题"`
	Detail    string    `xorm:"detail text" comment:"描述"`
	Status    int       `xorm:"status default(2)" comment:"处理状态 1: 已发送 2: 未处理"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
