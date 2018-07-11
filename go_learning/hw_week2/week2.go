package main

import (
	"fmt"
	"sort"
)

func main() {

	score := [5]int{49, 29, 12, 42, 52}

	var total int
	for i := 0; i < len(score); i++ {
		total = score[i] + total
	}
	fmt.Println("total => ", total)
	fmt.Println("AVG => ", total/len(score))

	findMiniNo()

}

func findMiniNo() {

	score := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	sort.Ints(score)

	fmt.Println("max => ", score[len(score)-1])

	fmt.Println("min => ", score[0])

}
