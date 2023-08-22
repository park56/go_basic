func solution(arr []int, query []int) []int {
    result := getQuery(arr, query)
    return result
}

func getQuery(arr []int, query []int ) []int {
    for i, v := range query {
        for _, v1 := range arr {
            switch isEven(i):
            true: 
				flag := 0
				for i, v := range arr {
					if v == query[queryNum] {
						flag = i
						break
					}
				}
				arr = arr[:flag]
            false: 
				flag := 0
				for i, v := range arr {
					if v == query[queryNum] {
						flag = i
						break
					}
				}
				arr = arr[flag:]
        }
    }
    return arr
}

// 짝수 판별
func isEven(num int) bool {
   return num % 2 == 0
}