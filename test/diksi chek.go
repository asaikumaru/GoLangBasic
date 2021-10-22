package main

import "fmt"

func main() {

	var chek = map[string]int{
		"turkey":      250,
		"turkeyFarsh": 140,
		"tomato":      150,
		"cheese":      100,
		"grechka":     70,
		"ovsanka":     14,
		"fassol":      100,
	}

	numsChek := make([]int, 0, len(chek))

	for _, value := range chek {
		numsChek = append(numsChek, value)
	}

	fmt.Println(calc(numsChek...))
}

func calc(nums ...int) int {
	result := 0

	for _, v := range nums {
		result += v
	}

	return result
}
