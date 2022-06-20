package main

import (
	"fmt"
	"reflect"
)

func main() {
	var m map[string]int
	var n map[string]int

	fmt.Println(m == nil)
	fmt.Println(n == nil)

	// 不能通过编译
	//fmt.Println(m == n)

	// reflect.DeepEqual() 可以比较两个 map 的内容是否相等
	fmt.Println(reflect.DeepEqual(m, n))
}

/*
map 深度相等的条件：
1、都为 nil
2、非空、长度相等，指向同一个 map 实体对象
3、相应的 key 指向的 value “深度”相等
*/
