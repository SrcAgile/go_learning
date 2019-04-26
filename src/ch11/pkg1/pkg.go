package news

import "fmt"

func init() {
	fmt.Println("the news pkg is loading successful")
}

func init() {
	fmt.Println("i am no.2")
}
func Spreading() {
	fmt.Println("bad news: you have been recived by THU")
}

// 1、对同一个go文件的init()调用顺序是从上到下的
// 2、对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数,对于
// 3、对不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()
// 4、如果package存在依赖，则先调用最早被依赖的package中的init()
