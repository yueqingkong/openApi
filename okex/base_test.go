package okex

import (
	"log"
	"testing"
)

func TestUsdCny(t *testing.T) {
	base := &Base{}
	rate := base.UsdCny()
	log.Println(rate)
}

