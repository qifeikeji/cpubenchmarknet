package cpubenchmarknet_test

import (
	"testing"

	"github.com/elliotwutingfeng/cpubenchmarknet"
	"github.com/perimeterx/marshmallow"
)

func TestGetCPUMegaList(t *testing.T) {
	CPUMegaList, err := cpubenchmarknet.GetCPUMegaList()
	if err != nil {
		t.Fatalf("Error downloading CPUMegaList, got: %v", err)
	}

	res, err := marshmallow.Unmarshal([]byte(CPUMegaList), &struct{}{})
	if err != nil {
		t.Fatalf("Error unmarshalling CPUMegaList, got: %v", err)
	}

	numKeys := len(res)
	if numKeys != 1 {
		t.Fatalf("Expected numKeys == 1, got %d", numKeys)
	}
	entries, ok := res["data"]
	if !ok {
		t.Fatal("Failed to access value of key 'data'")
	}
	numEntries := len(entries.([]interface{}))
	minimumNumEntries := 4800
	if numEntries < minimumNumEntries {
		t.Fatalf("Expected numEntries >= %d, got %d", minimumNumEntries, numEntries)
	}
}
