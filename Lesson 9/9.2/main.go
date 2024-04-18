package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	for _, v := range s {
		fmt.Println("v1:", v)
	out:
		for _, v2 := range s {
			fmt.Println("\tv2:", v2)
			for _, v3 := range s {
				fmt.Println("\t\tv3:", v3)
				for i, v4 := range s {
					fmt.Println("\t\t\tv4:", v4)
					if i == 1 {
						break out
					}
				}
			}
		}
	}
}
