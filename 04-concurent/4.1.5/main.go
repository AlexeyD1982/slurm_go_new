package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//waitGroupExample()
	//waitGroupExampleCopy()
	//mutexExample()
	//rwMutexExample()
	//syncMapExample()
	//raceExample()

	// blocking with timeout
	//tm := &TimeoutMutex{}
	//wg := &sync.WaitGroup{}
	//wg.Add(2)
	//go lock(tm, wg)
	//go lock(tm, wg)
	//wg.Wait()

	//phase
	var phase1WG, phase2WG sync.WaitGroup

	for i := 0; i < 3; i++ {
		phase1WG.Add(1)
		go phase(i, 1, &phase1WG)
	}

	phase1WG.Wait()
	fmt.Println("Phase 1 completed, starting Phase 2")

	for i := 0; i < 2; i++ {
		phase2WG.Add(1)
		go phase(i, 2, &phase2WG)
	}
	phase2WG.Wait()
	fmt.Println("All phases completed")
}

func process() {
	time.Sleep(2 * time.Second)
	fmt.Println("process")
}

func waitGroupExample() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		process()
	}()

	wg.Wait()
}

func waitGroupExampleCopy() {
	var wg sync.WaitGroup

	wg.Add(1)
	go processWithWg(&wg)

	wg.Wait()
}

func processWithWg(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("processWithWg")
}

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.m.Lock()
	defer c.m.Unlock()
	c.value += n
}

func mutexExample() {
	var wg sync.WaitGroup

	c := Counter{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go c.Update(10, &wg)
	}

	wg.Wait()
	fmt.Printf("Result is #%d\n", c.value)
}

type RWCounter struct {
	m     sync.RWMutex
	value int
}

func (c *RWCounter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.m.Lock()
	defer c.m.Unlock()
	c.value += n
}

func (c *RWCounter) GetValue() int {
	c.m.RLock()
	v := c.value
	defer c.m.RUnlock()

	return v
}

func rwMutexExample() {
	var wg sync.WaitGroup

	c := RWCounter{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go c.Update(i, &wg)
	}

	wg.Wait()
	fmt.Printf("Result is #%d\n", c.value)
}
func syncMapExample() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(10)

	for i := 1; i <= 5; i++ {
		go func(k int) {
			v := fmt.Sprintf("value #%d", k)

			fmt.Println("Writing: ", v)
			m.Store(k, v)
			wg.Done()
		}(i)
	}

	for i := 1; i <= 5; i++ {
		go func(k int) {
			v, _ := m.Load(k)
			fmt.Println("Reading: ", v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
func raceExample() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a"
		c <- true
	}()
	m["2"] = "b"
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// blocking with timeout

type TimeoutMutex struct {
	mu sync.Mutex
}

func (t *TimeoutMutex) LockWithTimeout(d time.Duration) bool {
	ch := make(chan struct{}, 1)
	go func() {
		t.mu.Lock()
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return true
	case <-time.After(d):
		return false
	}
}

func (t *TimeoutMutex) Unlock() {
	t.mu.Unlock()
}

func lock(tm *TimeoutMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	if tm.LockWithTimeout(1 * time.Second) {
		fmt.Println("Lock acquired")
		time.Sleep(2 * time.Second)
		tm.Unlock()
		fmt.Println("Lock released")
	} else {
		fmt.Println("Failed to acquire lock: timed out")
	}
}

// phase
func phase(id, phaseNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Phase %d, Task %d started\n", phaseNum, id)
	time.Sleep(time.Second)
	fmt.Printf("Phase %d, Task %d completed\n", phaseNum, id)
}
