package main

import (
	"fmt"
	
)

func reverse(s1 []int){
	var temp int
	len:=len(s1)

	for i:=0;i< len/2;i++{
		temp=s1[0+i]
		s1[0+i]=s1[ len-1-i]
		s1[len-1-i]=temp 
	}
	
}

func find_occ(str string) map[string]int {
	occ := make(map[string]int)
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		occ[char] += 1
	}
	return occ
}

func sum(sl []int) int {
	res := 0
	for i := 0; i < len(sl); i++ {
		res += sl[i]
	}
	return res
}

func calc_value_sum(sl_map map[int][]int) map[int]int {

	res := map[int]int{}
	for index, value := range sl_map {
		res[index] = sum(value)
	}

	return res
}

func convert_to_map(sli []int) map[int]int {

	res := map[int]int{}
	for i := 0; i < len(sli); i++ {
		res[i] = sli[i]
	}
	return res
}

func third() {

	s1 := []int{1, 2, 3, 4, 5}
	reverse(s1)
	fmt.Println("reverse slice is ",s1)

	str := "zopsmart interns"
	mp := find_occ(str)
	fmt.Println(str, " has the occurence of ", mp)

	sl_map := map[int][]int{
		1: []int{1, 2, 3, 4, 5},
		2: []int{1, 2, 3, 22, 2},
	}

	sum_map := calc_value_sum(sl_map)
	fmt.Println("the sum calculated value of the map is ", sum_map)

	slice_map := convert_to_map(s1)

	fmt.Println("the converted map of slice ", s1, " is ", slice_map)

}
