package qkl123

type Candle struct {
	Close       float32 `json:"close"`
	Date        int64   `json:"date"`
	FundsBigIn  float32 `json:"funds_big_in"`
	FundsBigOut float32 `json:"funds_big_out"`
	FundsIn     float32 `json:"funds_in"`
	FundsOut    float32 `json:"funds_out"`
	High        float32 `json:"high"`
	Low         float32 `json:"low"`
	Open        float32 `json:"open"`
	Vol         float32 `json:"vol"`
}
