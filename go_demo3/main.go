package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	int_list := []int{2, 34, 545, 1}
	slices.Sort(int_list)
	fmt.Println(int_list)

	slices.SortFunc(int_list, func(a, b int) int {
		return -cmp.Compare(a, b)
	})
	fmt.Println(int_list)

	// var userinfo map[string]string
	userinfo := make(map[string]string)
	userinfo["username"] = "张三"
	fmt.Println(userinfo)
}