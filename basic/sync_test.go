package basic

import (
	"testing"
)

func TestAutomicCounter(t *testing.T) {
	t.Log("Testing ExecAtomicCounter")
	size := int(100)
	count := int(100)
	var counter = new(AtomicCounter)
	counter.value = 0
	counter.waiter.Add(size)
	for i := 0; i < size; i++ {
		go ExecAtomicCounter(counter, count)
	}
	counter.waiter.Wait()
	if size*count != int(counter.value) {
		t.Fatalf("Counter: %d, Excepted %d", counter, size*count)
	}
}

func TestExecMutexCounter(t *testing.T) {
	t.Log("Testing ExecMutexCounter")
	size := int(100)
	count := int(100)
	var counter = new(MutexCounter)
	counter.value = 0
	counter.waiter.Add(size)
	for i := 0; i < size; i++ {
		go ExecMutexCounter(counter, count)
	}
	counter.waiter.Wait()
	if size*count != int(counter.value) {
		t.Fatalf("Counter: %d, Excepted %d", counter.value, size*count)
	}
}
