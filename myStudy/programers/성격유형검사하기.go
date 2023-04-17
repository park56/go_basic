package main

import (
	"sort"
)

var survey []int

func main() {

	r := 0
	t := 0
	f := 0
	c := 0
	m := 0
	j := 0
	a := 0
	n := 0

	elements := [8]string{"RT", "TR", "FC", "CF", "MJ", "JM", "AN", "NA"}

	sl := sort.StringSlice(elements[:])

	sl.Sort()

	for i, _ := range sl {
		for j := 0; i < 8; i++ {
			if sl.Search(sl[j]) != 0 {

				tmp := sl[sl.Search(sl[j])]

				if tmp == "RT" || tmp == "TR" {
					if survey[sl.Search(sl[j])+1] > 4 {
						r++
					}else if {
						
					}
				} else if tmp == "FC" || tmp == "CF" {

				} else if tmp == "MJ" || tmp == "JM" {

				} else {

				}
			}

		}
	}
}
