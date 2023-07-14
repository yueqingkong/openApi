package okex

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Ticker struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	Ts        string `json:"ts"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
}

type Instrument struct {
	InstType   string `json:"instType"`
	InstID     string `json:"instId"`
	InstFamily string `json:"instFamily"`
	Uly        string `json:"uly"`
	Category   string `json:"category"`
	BaseCcy    string `json:"baseCcy"`
	QuoteCcy   string `json:"quoteCcy"`
	SettleCcy  string `json:"settleCcy"`
	CtVal      string `json:"ctVal"`
	CtMult     string `json:"ctMult"`
	CtValCcy   string `json:"ctValCcy"`
	OptType    string `json:"optType"`
	Stk        string `json:"stk"`
	ListTime   string `json:"listTime"`
	ExpTime    string `json:"expTime"`
	Lever      string `json:"lever"`
	TickSz     string `json:"tickSz"`
	LotSz      string `json:"lotSz"`
	MinSz      string `json:"minSz"`
	CtType     string `json:"ctType"`
	Alias      string `json:"alias"`
	State      string `json:"state"`
}

type PriceLimit struct {
	InstType string `json:"instType"`
	InstID   string `json:"instId"`
	BuyLmt   string `json:"buyLmt"`
	SellLmt  string `json:"sellLmt"`
	Ts       string `json:"ts"`
}

type EstimatedPrice struct {
	InstType string `json:"instType"`
	InstID   string `json:"instId"`
	SettlePx string `json:"settlePx"`
	Ts       string `json:"ts"`
}

type SysTime struct {
	Ts string `json:"ts"`
}

type Candle [][]string

type OrderRes struct {
	ClOrdID string `json:"clOrdId"`
	OrdID   string `json:"ordId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

type Balance struct {
	AdjEq   string `json:"adjEq"`
	Details []struct {
		AvailBal      string `json:"availBal"`
		AvailEq       string `json:"availEq"`
		CashBal       string `json:"cashBal"`
		Ccy           string `json:"ccy"`
		CrossLiab     string `json:"crossLiab"`
		DisEq         string `json:"disEq"`
		Eq            string `json:"eq"`
		EqUsd         string `json:"eqUsd"`
		FrozenBal     string `json:"frozenBal"`
		Interest      string `json:"interest"`
		IsoEq         string `json:"isoEq"`
		IsoLiab       string `json:"isoLiab"`
		IsoUpl        string `json:"isoUpl"`
		Liab          string `json:"liab"`
		MaxLoan       string `json:"maxLoan"`
		MgnRatio      string `json:"mgnRatio"`
		NotionalLever string `json:"notionalLever"`
		OrdFrozen     string `json:"ordFrozen"`
		Twap          string `json:"twap"`
		UTime         string `json:"uTime"`
		Upl           string `json:"upl"`
		UplLiab       string `json:"uplLiab"`
		StgyEq        string `json:"stgyEq"`
	} `json:"details"`
	Imr         string `json:"imr"`
	IsoEq       string `json:"isoEq"`
	MgnRatio    string `json:"mgnRatio"`
	Mmr         string `json:"mmr"`
	NotionalUsd string `json:"notionalUsd"`
	OrdFroz     string `json:"ordFroz"`
	TotalEq     string `json:"totalEq"`
	UTime       string `json:"uTime"`
}

type OrderInfo struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstType           string `json:"instType"`
		InstId             string `json:"instId"`
		Ccy                string `json:"ccy"`
		OrdId              string `json:"ordId"`
		ClOrdId            string `json:"clOrdId"`
		Tag                string `json:"tag"`
		Px                 string `json:"px"`
		Sz                 string `json:"sz"`
		Pnl                string `json:"pnl"`
		OrdType            string `json:"ordType"`
		Side               string `json:"side"`
		PosSide            string `json:"posSide"`
		TdMode             string `json:"tdMode"`
		AccFillSz          string `json:"accFillSz"`
		FillPx             string `json:"fillPx"`
		TradeId            string `json:"tradeId"`
		FillSz             string `json:"fillSz"`
		FillTime           string `json:"fillTime"`
		Source             string `json:"source"`
		State              string `json:"state"`
		AvgPx              string `json:"avgPx"`
		Lever              string `json:"lever"`
		AttachAlgoClOrdId  string `json:"attachAlgoClOrdId"`
		TpTriggerPx        string `json:"tpTriggerPx"`
		TpTriggerPxType    string `json:"tpTriggerPxType"`
		TpOrdPx            string `json:"tpOrdPx"`
		SlTriggerPx        string `json:"slTriggerPx"`
		SlTriggerPxType    string `json:"slTriggerPxType"`
		SlOrdPx            string `json:"slOrdPx"`
		StpId              string `json:"stpId"`
		StpMode            string `json:"stpMode"`
		FeeCcy             string `json:"feeCcy"`
		Fee                string `json:"fee"`
		RebateCcy          string `json:"rebateCcy"`
		Rebate             string `json:"rebate"`
		TgtCcy             string `json:"tgtCcy"`
		Category           string `json:"category"`
		ReduceOnly         string `json:"reduceOnly"`
		CancelSource       string `json:"cancelSource"`
		CancelSourceReason string `json:"cancelSourceReason"`
		QuickMgnType       string `json:"quickMgnType"`
		AlgoClOrdId        string `json:"algoClOrdId"`
		AlgoId             string `json:"algoId"`
		UTime              string `json:"uTime"`
		CTime              string `json:"cTime"`
	} `json:"data"`
}

type SetLeverage struct {
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	InstID  string `json:"instId"`
	PosSide string `json:"posSide"`
}

type UsdCny struct {
	UsdCny string `json:"usdCny"`
}
