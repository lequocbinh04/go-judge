package appCommon

import (
	"math/big"
	"strings"
)

func ConvertTopicToBigInt(topic string) *big.Int {
	ii := new(big.Int)
	ii.SetString(strings.ReplaceAll(topic, "0x", ""), 16)
	return ii
}
