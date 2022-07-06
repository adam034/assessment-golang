package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n = []int{1, 4, 5, 6, 8, 2}
	var i = 1
	for i < len(n) {
		var j = i
		for j >= 1 && n[j] > n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]

			j--
		}

		i++
		for i := 8; i > 0; i-- {
			var bar []string
			for j := 0; j < len(n); j++ {
				if n[j] < i {
					bar = append(bar, "")
				} else {
					bar = append(bar, "|")
				}

			}
			fmt.Println(strings.Join(bar, "\t"))
		}
		var nums []string
		for _, i := range n {
			nums = append(nums, strconv.Itoa(i))
		}
		fmt.Println(strings.Join(nums, "\t"))
	}
}
