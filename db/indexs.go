package db

import (
	"errors"
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/util"
	"log"
	"time"
)

type Indexs struct {
	Id        int64
	Plat      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 平台名称
	Name      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 指标名称
	Symbol    string    `xorm:"varchar(255) unique(p-n-s-d)"` // 币种
	Date      string    `xorm:"varchar(255) unique(p-n-s-d)"` // 格式化时间
	P1        float32   `xorm:"p_1"`
	P2        float32   `xorm:"p_2"`
	P3        float32   `xorm:"p_3"`
	P4        float32   `xorm:"p_4"`
	P5        float32   `xorm:"p_5"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (self *Indexs) Get(pt conset.PLAT, name string, bs conset.CCY, quote conset.CCY, start time.Time) (*Indexs, error) {
	indexs := &Indexs{Plat: Plat(pt), Name: name, Symbol: Symbol(bs, quote), Date: util.TimeFormatDay(start)}
	if b, err := Engine().Get(indexs); err != nil {
		return nil, err
	} else if !b {
		return nil, errors.New("not exist")
	}

	return indexs, nil
}

func (self *Indexs) IndexLast(name string, bs, quote conset.CCY) (*Indexs, error) {
	indexs := &Indexs{Name: name, Symbol: Symbol(bs, quote)}
	if b, err := Engine().Desc("date").Get(indexs); err != nil {
		return nil, err
	} else if !b {
		return nil, errors.New("not exist")
	}

	return indexs, nil
}

func (self *Indexs) Create(pt conset.PLAT, name string, bs conset.CCY, quote conset.CCY, start time.Time, short, long, low, high, atr float32) error {
	record := &Indexs{Plat: Plat(pt), Name: name, Symbol: Symbol(bs, quote), Date: util.TimeFormatDay(start),
		P1: short, P2: long, P3: low, P4: high, P5: atr}

	_, err := Engine().InsertOne(record)
	return err
}

func (self *Indexs) IndexGetCreate(pt conset.PLAT, name string, bs conset.CCY, quote conset.CCY, times conset.TIMES, start time.Time, fc func() (float32, float32, float32, float32, float32)) (float32, float32, float32, float32, float32) {
	coin := &Coin{}
	lastCoin, _ := coin.Last(pt, bs, quote, times)
	indexs, err := self.IndexLast(name, bs, quote)

	var timeDif int32
	if times == conset.H_4 {
		timeDif = 4
	} else if times == conset.H_6 {
		timeDif = 6
	} else if times == conset.H_12 {
		timeDif = 12
	} else if times == conset.D_1 {
		timeDif = 24
	}

	// lastCoin.CreateTime 会比 indexs.Date 慢一个周期
	if err != nil || lastCoin.CreateTime.Add(time.Duration(timeDif)*time.Hour).Sub(util.StringToTime(indexs.Date)) > time.Duration(10)*time.Second {
		s, l, low, high, atr := fc()

		indexDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
		more := int32(start.Sub(indexDate).Hours())
		integerHour := more - more%timeDif
		indexDate = indexDate.Add(time.Duration(integerHour) * time.Hour)

		if err = self.Create(pt, name, bs, quote, indexDate, s, l, low, high, atr); err != nil {
			log.Printf("%s  IndexCreate err: %v", name, err)
		}

		log.Printf("%s  IndexCreate s: %v l: %v low: %v high: %v atr: %v", name, s, l, low, high, atr)
		return s, l, low, high, atr
	}

	return indexs.P1, indexs.P2, indexs.P3, indexs.P4, indexs.P5
}
