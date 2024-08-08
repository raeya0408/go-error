package main


import (
	"fmt"
	"strconv"
)

func main() {
	ints := []int{3, 7, 9, 4, 4}
	//切片int convert string
	fmt.Printf("%T", convert1(ints))
	fmt.Printf("%T", convert2(ints))
	fmt.Printf("%T", convert3(ints))

}

// 21实现一个convert方法，将切片Foo转换成另一个类型的切片Bar
// 法一
func convert1(ints []int) []string {
	strs := make([]string, 0)
	for _, s := range ints {
		strs = append(strs, strconv.Itoa(s))
	}
	return strs
}
// 法二
func convert2(ints []int) []string {
	n := len(ints)
	strs := make([]string, 0, n)
	for _, s := range ints {
		strs = append(strs, strconv.Itoa(s))
	}
	return strs
}
// 法三
func convert3(ints []int) []string {
	n := len(ints)
	strs := make([]string, n, n)
	for i, ints := range ints {
		strs[i] = strconv.Itoa(ints)
	}
	return strs
}

