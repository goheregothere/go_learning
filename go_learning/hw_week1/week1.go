package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a = 1
	var b int32 = 2
	var c int64 = 3
	var d = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x = "i love golang"

	var q1 = a + int(b)
	var q2 = a + int(b) + int(c)
	var q3 = float32(e) / float32(f)

	i, _ := strconv.Atoi(d)
	var q4 = a + i

	j := strconv.Itoa(a)

	fmt.Println("q1 => ", q1)
	fmt.Println("q2 => ", q2)
	fmt.Println("q3 => ", q3)
	fmt.Println("q4 => ", q4)
	fmt.Println("q5 => ", x, j)

	iAmPrivate()

}

func iAmPrivate() {
	fmt.Println("i am private func")
}
