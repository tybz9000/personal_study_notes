package main

//特殊语法1，const连续声明
const (
	a1 = 10
	a2 = iota
	a3
	count
)

func main(){

	println(a1)
	println(a2)
	println(a3)
	println(count)
}