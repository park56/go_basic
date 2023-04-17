func solution(nums []int) int {

    var temp []int
    
    for _,v1 := range nums {
        for _,v2 := range nums {
            for _,v3 := range nums {
                if v1 == v2 || v2 == v3 || v3 == v1{
                    continue
                }
                if checkMinNum(v1+v2+v3){
                    if !sameCheck(temp,v1+v2+v3){
                        temp = append(temp,v1+v2+v3)
                    }
                }
                continue
            }
        }
    }
    
    fmt.Println(temp)

    return len(temp)
}

func checkMinNum(answer int)bool{
    for i := 2; i<int(float64(answer)*0.5); i++ {
        if answer % i == 0{
            return false;
        }
    }
    return true;
}

func sameCheck(arr []int, num int)bool{
    for _,v := range arr {
        if v == num{
            return true
        }
    }
    return false
}
//=============================

9


123456789101112131415161718