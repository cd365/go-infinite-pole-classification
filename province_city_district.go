package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Rows []*Address = []*Address{
	&Address{
		Id:1,
		Name:"1",
		Parent:0,
	},
	&Address{
		Id:2,
		Name:"2",
		Parent:0,
	},
	&Address{
		Id:3,
		Name:"3",
		Parent:0,
	},
	&Address{
		Id:4,
		Name:"1-1",
		Parent:1,
	},
	&Address{
		Id:5,
		Name:"1-3",
		Parent:1,
	},
	&Address{
		Id:6,
		Name:"1-2",
		Parent:1,
	},
	&Address{
		Id:7,
		Name:"2-1",
		Parent:2,
	},
	&Address{
		Id:8,
		Name:"2-2",
		Parent:2,
	},
	&Address{
		Id:9,
		Name:"2-3",
		Parent:2,
	},
	&Address{
		Id:10,
		Name:"10",
		Parent:0,
	},
}

// Address
type Address struct {
	Id     int64      `json:"id"`
	Name   string     `json:"name"`
	Parent int64      `json:"parent"`
	Child  []*Address `json:"child,omitempty"`
}

// Infinite 无限极分类实现(无排序)
func Infinite(rows []*Address) []*Address {
	maps := make(map[int64]*Address)
	for _, v := range rows {
		maps[v.Id] = v
	}
	for _, v := range maps {
		if _, ok := maps[v.Parent]; ok {
			maps[v.Parent].Child = append(maps[v.Parent].Child, v)
		}
	}
	result := []*Address{}
	for _, v := range maps {
		if v.Parent == 0 {
			result = append(result, v)
		}
	}
	return result
}

// AddressInfiniteBubbleAsc 冒泡升序
func AddressInfiniteBubbleAsc(addr []*Address) []*Address {
	count := len(addr)
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if addr[i].Id > addr[j].Id {
				addr[i], addr[j] = addr[j], addr[i]
			}
		}
	}
	for i := 0; i < count; i++ {
		addr[i].Child = AddressInfiniteBubbleAsc(addr[i].Child)
	}
	return addr
}

// AddressInfiniteBubbleDesc 冒泡降序
func AddressInfiniteBubbleDesc(addr []*Address) []*Address {
	count := len(addr)
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if addr[i].Id < addr[j].Id {
				addr[i], addr[j] = addr[j], addr[i]
			}
		}
	}
	for i := 0; i < count; i++ {
		addr[i].Child = AddressInfiniteBubbleDesc(addr[i].Child)
	}
	return addr
}

func main() {
	rows := Rows
	start := time.Now()
	result := Infinite(rows)
	result = AddressInfiniteBubbleAsc(result)
	use := time.Now().Sub(start)
	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
	fmt.Println(use) // 973.6µs
	fmt.Println(len(result))
}
