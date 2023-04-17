func solution(numbers []int) int {

	res := 0
	oneToNine := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range numbers {
		for _, v2 := range oneToNine {
			if v == v2 {
				res += v
			}
		}
	}

	return 45 - res
}

// ======================= 더 나은 답
func solition(numbers []int) int {
	res := 0

	for _, v := range numbers {
		res += v
	}
	return 45 - res
}