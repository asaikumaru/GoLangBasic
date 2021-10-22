package main

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

	numsChek := []int{}

	for i, v := range numsChek {
		res := 0
		for a, v2 := range chek {

		}
		res += v
		return res

	}

	value := calc(chek...)
}

func calc(nums ...[]int) int {
	result := 0

	for _, v := range nums {
		for _, v2 := range v {
			result += v2
		}
	}

	return result
}
