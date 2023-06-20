package cpubenchmarknet_test

import (
	"fmt"

	"github.com/elliotwutingfeng/cpubenchmarknet"
)

func Example_basic() {
	CPUMegaList, err := cpubenchmarknet.GetCPUMegaList()
	if err == nil {
		fmt.Println(CPUMegaList) // JSON string
	} else {
		fmt.Println(err)
	}
}
