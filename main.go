package main


import (
	"fmt"
	"strconv"
)

func main() {
	ints := []int{3, 7, 9, 4, 4}

	//切片int convert string
	fmt.Printf("%T", convert1(ints))
	fmt.Println(convert1(ints))

	fmt.Printf("%T", convert2(ints))
	fmt.Println(convert2(ints))


	fmt.Printf("%T", convert22(ints))
	fmt.Println(convert22(ints))


	fmt.Printf("%T", convert3(ints))
	fmt.Println( convert3(ints))

	fmt.Printf("%T", convert33(ints))
	fmt.Println( convert33(ints))



}

// 21实现一个convert方法，将切片int转换成另一个类型的切片string
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
//若需要将一个int切片转换成一个两倍长的string切片，则使用索引复制的方式看起来不清晰，且不易维护
//若需要将一个Foo切片转换成一个两倍长的Bar切片，则使用索引复制的方式看起来不清晰，且不易维护
//法二
func convert22(ints []int)[]string{
	n:=len(ints)
	strs:=make([]string,0,2*n)
	for _,s:=range ints{
		strs=append(strs, strconv.Itoa(s))
		strs=append(strs, strconv.Itoa(s))
	}
	return strs
}
//法三
func convert33(ints []int)[]string{
	n := len(ints)
	strs := make([]string, 2*n)
	for i, ints := range ints {
		strs[2*i] = strconv.Itoa(ints)
		strs[2*i+1] = strconv.Itoa(ints)

	}
	return strs
}
