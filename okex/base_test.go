package okex

import (
	"github.com/yueqingkong/openApi/conset"
	"log"
	"testing"
)

// go test -v -run TestUsdCny
func TestUsdCny(t *testing.T) {
	base := &Base{}
	rate := base.UsdCny()
	log.Println(rate)
}

// go test -v -run TestOrders
func TestOrders(t *testing.T) {
	base := &Base{}
	base.Init([]string{"","",""})
	if b := base.Orders(conset.FIL_USDT, conset.SWAP, conset.BUY_HIGH, 6.378, 1); !b {
		t.Log(b)
	}
}