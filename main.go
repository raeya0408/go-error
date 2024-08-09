package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	err21()
	err22()
	err24()
	err25()





}

// 21实现一个convert方法，将切片int转换成另一个类型的切片string
func err21(){
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



//22切片为nil和空混淆
//nil,切片等于nil
//空，长度为零
//nil也是空切片，空切片不一定是nil切片
func err22(){
	var s1 []string//法一
	s2:=[]string(nil)//法二
	s3:=[]string{}//法三
	s4:=make([]string, 0)//法四
	
	fmt.Printf("empty=%t\tnil=%t\n",len(s1)==0,s1==nil)
	fmt.Printf("empty=%t\tnil=%t\n",len(s2)==0,s2==nil)
	fmt.Printf("empty=%t\tnil=%t\n",len(s3)==0,s3==nil)
	fmt.Printf("empty=%t\tnil=%t\n",len(s4)==0,s4==nil)	
}



//23没有正确检查切片是否为空
//切片用法三法四初始化
//错误写法
// func handleOperations(id string){
// 	operations:=getOperation(id)
// 	if operations!=nil{
// 		handle(operations)
// 	}
// }
// func getOperation(id string)[]string{
// 	operation:=make([]string,0)
// 	if id==""{
// 		return operation
// 	}
// 	return operation
// }
//正确写法
//法一	修改被调用方
// func getOperation(id string)[]string{
// 	operation:=make([]string,0)
// 	if id==""{
// 		return nil
// 	}
// 	return operation
// }
//法二	修改调用方
// func handleOperations(id string){
// 	operations:=getOperation(id)
// 	if len(operations)!=0{
// 		handle(operations)
// 	}
// }



//24错误的切片拷贝
//错误示例
// src:=[]int{1,2,3}
// var dst []int
// copy(dst,src)//原因在于内置的copy函数，拷贝的切片的元素个数等于min（len（电视台），len（双人床））
// fmt.Println(dst)
func err24(){
	src:=[]int{1,2,3}
	//法一
	dst1:=make([]int,len(src))
	copy(dst1,src)
	fmt.Println(dst1)
	//法二
	dst2:=append([]int(nil),src...)
	fmt.Println(dst2)
}



//25切片使用append的副作用
// func 示例(){
// 	//示例
// 	s1:=[]int{1,2,3}
// 	s2:=s1[1:2]
// 	s3:=append(s2,10)
// 	fmt.Println(s1,s2,s3)//s1切片第三个元素也发生了修改

// 	s:=[]int{6,7,8}
// 	f(s[:2])//亦发生在切片作为参数传递给某个函数
// 	fmt.Println(s)
// }
func f(s []int){
	_=append(s,10)
}
func err25(){
	//法一
	s:=[]int{1,2,3}
	sCopy:=make([]int,2)
	copy(sCopy,s)
	f(sCopy)
	fmt.Println(s)
	//法二
	f(s[:2:2])//第三个参数规定cap
}



//26切片和内存泄漏
//场景一
// func consumeMessage(){
// 	for{
// 		msg:=receiveMessage()//假设每次msg都是一个长度为100000的字节切片
// 		storeMessageType(getMessageType(msg))
// 	}
// }
//字节切片截取函数，截取前五个字符
// func getMessageType(msg []byte)[]byte{
// 	return msg[:5]
// }

//解决方案
//有效
func getMessageType(msg []byte)[]byte{
	msgType:=make([]byte,5)
	copy(msgType,msg)
	return msgType
}
//无效
// func getMessageType(msg []byte)[]byte{
// 	return msg[:5:5]
// }
//场景二切片和引用
type S struct{
	s []byte
}
func SS(){
	Ss:=make([]S,1_000)
	printAlloc()
	for i:=0;i<len(Ss);i++{
		Ss[i]=S{
			s:make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two:=keepFirstTwoElementsOnly(Ss)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)//保持对变量two的引用
}
func keepFirstTwoElementsOnly(Ss []S)[]S{
	return Ss[:2]
}
func printAlloc(){}


//27