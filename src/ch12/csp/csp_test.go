package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func Myservice() string {
	fmt.Println("call my service")
	time.Sleep(time.Millisecond * 300) //suppose io operation
	return "return my service"
}

func AsyncMyservice() chan string {
	channel := make(chan string)
	go func() {
		ret := Myservice()

		channel <- ret
		fmt.Println("return my go func")
	}()

	return channel
}

func otherTask() {
	time.Sleep(time.Millisecond * 10)
	fmt.Println("return other Task")
}

func TestChannel(t *testing.T) {
	channel := AsyncMyservice()
	otherTask()
	fmt.Println(<-channel)
}
