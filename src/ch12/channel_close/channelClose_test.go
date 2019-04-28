package channelClose

import (
	"fmt"
	"sync"
	"testing"
)

func dataSender(channel chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10000; i++ {
			channel <- i
		}
		close(channel)
		wg.Done()
	}()

}

func dataReciever(channel chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-channel; ok {
				fmt.Println("the data is", data)
			} else {
				break
			}

		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup

	channel := make(chan int)

	wg.Add(1)
	dataSender(channel, &wg)
	wg.Add(1)
	dataReciever(channel, &wg)

	wg.Wait()
}
