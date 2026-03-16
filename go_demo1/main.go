package main

import (
	"fmt"
)

func main() {
	var intArr1 = []int{1, 4, 323, 435, 6}
	sum := 0
	for idx, key := range intArr1 {
		fmt.Println(idx)
		sum += key
	}
	for i := 0; i < len(intArr1); i++ {
		sum += intArr1[i]
	}
	fmt.Println(sum, sum/(2*len(intArr1)))
	fmt.Println(sum, float32(sum)/float32(2*len((intArr1))))

	var intArr2 = [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sum_2 := 0
	for idx, key1 := range intArr2 {
		fmt.Println(idx)
		fmt.Println(key1)
		for _, key2 := range key1 {
			sum_2 += key2
			fmt.Println(key2)
		}
	}

	for i := 0; i < len(intArr2); i++ {
		for j := 0; j < len(intArr2[i]); j++ {
			fmt.Print(intArr2[i][j])
		}
	}
}
