package util

import (
	"testing"
)

// go test -v -run TestPayFee
func TestPayFee(t *testing.T) {
	t.Log("fee: ", PayFee(10000, 100, 100))
}
