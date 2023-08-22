package main

func solution(babbling []string) int {
	canSpeak := []string{"aya", "ye", "woo", "ma"}
	canSpeakList := getSpeaks(canSpeak)
	canSpeakList = append(canSpeakList, &canSpeak)
	result := lastDance(babbling, canSpeakList)
	return result
}

func getSpeaks(list []string) []string {
	newList := []string{}
	for _, v := range list {

		for _, v2 := range list {
			if v == v2 {
				continue
			}
			newList = append(newList, v+v2)

			for _, v3 := range list {
				if v == v2 || v == v3 || v2 == v3 {
					continue
				}
				newList = append(newList, v+v2+v3)

				for _, v4 := range list {
					if v == v2 || v == v3 || v2 == v3 || v == v4 || v2 == v4 || v3 == v4 {
						continue
					}
					newList = append(newList, v+v2+v3+v4)
				}
			}
		}
	}
	return newList
}

func lastDance(bab []string, list []string) int {
	flag := 0
	for _, v := range bab {
		for _, v2 := range list {
			if v == v2 {
				flag++
			}
		}
	}

	return flag
}
