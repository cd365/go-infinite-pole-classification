package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func BenchmarkAddressInfiniteBubbleAsc(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		AddressInfiniteBubbleAsc(Infinite(Rows))
	}
}

func TestAddressInfiniteBubbleAsc(t *testing.T) {
	result := AddressInfiniteBubbleAsc(Infinite(Rows))
	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
}
