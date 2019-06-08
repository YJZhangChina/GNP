package basic

import (
	"fmt"
	"sync"
	"testing"
)

func TestUnsafePointerPassing(t *testing.T) {
	fmt.Println("Testing UnsafePointerPassing in GoroutineTraps")
	size := 10
	var waiter sync.WaitGroup
	waiter.Add(size)
	message := &Message{current: 1}
	for i := 0; i < size; i++ {
		go valueChange(i+1, message, &waiter)
	}
	waiter.Wait()
}

func TestSafePointerPassing(t *testing.T) {
	fmt.Println("Testing SafePointerPassing in GoroutineTraps")
	size := 10
	var waiter sync.WaitGroup
	waiter.Add(size)
	message := &Message{current: 1}
	for i := 0; i < 10; i++ {
		var copy = new(Message)
		*copy = *message
		go valueChange(i+1, copy, &waiter)
	}
	waiter.Wait()
}
