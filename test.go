package main

import "fmt"

func myFunc() {
	c := make(chan int, 1)
	c <- 1 // 报错 deadlock, 下面的代码不会执行
	fmt.Println(<-c)
}

func main() {
	myFunc()
}
