package main

import "fmt"

func main() {
	ageMap := make(map[string]int)
	ageMap["qcrao"] = 18

	// 不带 comma 用法
	age1 := ageMap["stefno"]
	fmt.Println(age1)

	// 带 comma 用法
	age2, ok := ageMap["stefno"]
	fmt.Println(age2, ok)
}

/*
注意：
在背后的编译，针对这两种情况，底层会调用两个不同的函数：
// src/runtime/hashmap.go
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer
func mapaccess2(t *maptype, h *hmap, key unsafe.Pointer) (unsafe.Pointer, bool)
两者的代码也是完全一样的，mapaccess2 函数返回值多了一个 bool 型变。
*/
