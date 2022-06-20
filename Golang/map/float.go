package main

import (
	"fmt"
	"math"
)

func main() {
	m := make(map[float64]int)
	m[1.4] = 1
	m[2.4] = 2
	m[math.NaN()] = 3
	m[math.NaN()] = 3

	for k, v := range m {
		fmt.Printf("[%v, %d] ", k, v)
	}

	fmt.Printf("\nk: %v, v: %d\n", math.NaN(), m[math.NaN()])
	fmt.Printf("k: %v, v: %d\n", 2.400000000001, m[2.400000000001])
	fmt.Printf("k: %v, v: %d\n", 2.4000000000000000000000001, m[2.4000000000000000000000001])

	fmt.Println(math.NaN() == math.NaN())
}

/*
当用 float64 作为 key 的时候，先要将其转成 uint64 类型，再插入 key 中。
float 型是可以作为 key 的，但是由于精度的问题，会导致一些诡异的问题，慎用之

// NAN（not a number）
// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float64 { return Float64frombits(uvnan) }
NAN 的特性：
1，NAN != NAN
2，hash(NAN) != hash(NAN)
多个 NAN 在插入 map 时，会被认为是不同的 key；

当 key 是引用类型时，判断两个 key 是否相等，需要 hash 后的值相等并且 key 的字面量相等
*/

func Float64bits() {
	m := make(map[float64]int)
	m[2.4] = 2

	fmt.Println(math.Float64bits(2.4))
	fmt.Println(math.Float64bits(2.400000000001))
	fmt.Println(math.Float64bits(2.4000000000000000000000001))
}
