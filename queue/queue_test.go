package queue

import (
	"fmt"
	"testing"
)

func TestQueue(test *testing.T) {
	forever := make(chan struct{})
	//today := time.Now().Format("20060102")
	queue := newQueue()
	queue.Run() // producer


	fmt.Println(queue.Progress())

	<-forever

}
