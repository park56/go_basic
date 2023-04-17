func solution(absolutes []int, signs []bool) int {

	res := 0

	for i, v := range signs {
		if v {
			res += absolutes[i]
		} else {
			res -= absolutes[i]
		}
	}
	return res
}