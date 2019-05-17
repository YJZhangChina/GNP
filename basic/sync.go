package basic

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value  int64
	waiter sync.WaitGroup
}

type MutexCounter struct {
	value  int64
	waiter sync.WaitGroup
	mutex  sync.Mutex
}

func ExecAtomicCounter(counter *AtomicCounter, count int) {
	defer counter.waiter.Done()
	for i := 0; i < count; i++ {
		atomic.AddInt64(&counter.value, 1)
		runtime.Gosched()
	}
}

func ExecMutexCounter(counter *MutexCounter, count int) {
	defer counter.waiter.Done()
	for i := 0; i < count; i++ {
		counter.mutex.Lock()
		{
			value := counter.value
			runtime.Gosched()
			value++
			counter.value = value
		}
		counter.mutex.Unlock()
	}
}
