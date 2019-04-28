package corotine_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestCoroutine(t *testing.T) {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 10000; i++ {

		wg.Add(1)
		go func() {
			defer mutex.Unlock()
			mutex.Lock()
			count = count + 1
			fmt.Println("the count is", count)
			wg.Done()
		}()

	}
	wg.Wait()
	//防止主线程提前退出
	//time.Sleep(1 * time.Second)
}
