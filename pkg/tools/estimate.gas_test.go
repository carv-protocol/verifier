package tools

import (
	"fmt"
	"testing"
)

func TestEstimateGas(t *testing.T) {
	gas := EstimateGas("https://arb-mainnet.g.alchemy.com/v2/xx", "0x4681Cf59A932ff22d8c9d36c17c28c2ae23f9EC5", "0xFFbB58c8370f99b2ae619328D6B99D77Fef190Cb", "NodeExit()", 0)
	fmt.Println(gas)
}
