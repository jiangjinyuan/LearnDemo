package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},len: 10,cap: 10
	s1 := slice[2:5]
	// s1 从 slice 索引2（闭区间）到索引5（开区间，元素真正取到索引4），长度为3，容量默认到数组结尾，为8。
	// s1: []int{2, 3, 4},len: 3,cap: 8
	s2 := s1[2:6:7]
	// s2 从 s1 的索引2（闭区间）到索引6（开区间，元素真正取到索引5），容量到索引7（开区间，真正到索引6），为5。
	// s2: []int{4, 5, 6, 7},len: 4,cap: 5

	s2 = append(s2, 100)
	// 向 s2 尾部追加一个元素 100，s2 容量刚好够，直接追加，不会扩容，但这会修改原始数组对应位置的元素，此时会影响到原来的 s1 和 slice。
	// s2: []int{4, 5, 6, 7, 100},len: 5,cap: 5
	// s1: []int{2, 3, 4, 5, , , 100, },len: 3,cap: 8
	// slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 100, 9},len: 10,cap: 10
	s2 = append(s2, 200)
	// 再向 s2 尾部追加一个元素 200，s2 容量不够，会扩容，s2 另起炉灶，将原来的元素复制新的位置，扩大自己的容量，此时不会影响到原来的 s1 和 slice。
	// s2: []int{4, 5, 6, 7, 100, 200},len: 6,cap: 10
	// s1: []int{2, 3, 4, 5, , , 100, },len: 3,cap: 8
	// slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 100, 9},len: 10,cap: 10

	s1[2] = 20
	// 将 s1 的索引2的元素修改为 20，此时不会影响到 s2, 会影响到原来的slice。
	// s1: []int{2, 3, 20, 5, , , 100, },len: 3,cap: 8
	// slice: []int{0, 1, 2, 3, 20, 5, 6, 7, 100, 9},len: 10,cap: 10

	// [2 3 20]
	fmt.Println(s1)
	// [4 5 6 7 100 200]
	fmt.Println(s2)
	// [0 1 2 3 20 5 6 7 100 9]
	fmt.Println(slice)
}

/*
注意点：
1）slice 的底层数据是数组，slice 是对数组的封装，它描述一个数组的片段；
2）底层数组是可以被多个 slice 同时指向的，因此对一个 slice 的元素进行操作是有可能影响到其他 slice 的元素的；
3）在打印slice的时候，只会打印长度以内的数据，即使底层数组还有其它的数据，也不会打印；
*/
