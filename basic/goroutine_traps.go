package basic

import (
	"fmt"
	"sync"
)

type Message struct {
	current int
}

func valueChange(step int, message *Message, waiter *sync.WaitGroup) {
	message.current = step
	fmt.Printf("After step %d the value is %d.\n", step, message.current)
	waiter.Done()
}
