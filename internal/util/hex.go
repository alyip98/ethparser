package util

import (
	"math/big"
	"strings"
)

func HexToInt64(hex string) int64 {
	i := big.Int{}
	i.SetString(strings.TrimPrefix(hex, "0x"), 16)
	return i.Int64()
}

func Int64ToHex(x int64) string {
	i := big.NewInt(x)
	return "0x" + i.Text(16)
}
