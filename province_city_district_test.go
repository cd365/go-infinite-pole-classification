package main

import (
	"testing"
)

func TestInfinite(b *testing.T) {
	// test empty data
	rows := []*Address{}
	Infinite(rows)

	// not empty data
	Infinite(Rows)
}

func TestAddressInfiniteBubbleAsc(t *testing.T) {
	AddressInfiniteBubbleAsc(Infinite(Rows))
}

func BenchmarkAddressInfiniteBubbleAsc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddressInfiniteBubbleAsc(Infinite(Rows))
	}
}
