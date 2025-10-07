package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap wraps a map with RWMutex for safe concurrent access
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func (s *SafeMap) Get(id, key string) int {
	fmt.Printf("Reader %s: trying RLock\n", id)
	s.mu.RLock()
	fmt.Printf("Reader %s: acquired RLock\n", id)

	val := s.m[key]

	s.mu.RUnlock()
	fmt.Printf("Reader %s: released RLock\n", id)
	return val
}

func (s *SafeMap) Set(key string, value int) {
	fmt.Println("Writer: trying Lock")
	s.mu.Lock()
	fmt.Println("Writer: acquired Lock")

	s.m[key] = value

	s.mu.Unlock()
	fmt.Println("Writer: released Lock")
}

func main() {
	safe := &SafeMap{
		m: map[string]int{"x": 0, "y": 0},
	}

	// Stop everything after 5 seconds
	done := time.After(5 * time.Second)

	// Writer goroutine (updates every second)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				time.Sleep(1 * time.Second)
				x := safe.Get("writer", "x") + 1
				y := safe.Get("writer", "y") + 1
				safe.Set("x", x)
				safe.Set("y", y)
			}
		}
	}()

	// 3 reader goroutines for demo (reduce from 10 to avoid spam)
	for i := 1; i <= 3; i++ {
		id := fmt.Sprintf("%d", i)
		go func(id string) {
			for {
				select {
				case <-done:
					return
				default:
					_ = safe.Get(id, "x")
					_ = safe.Get(id, "y")
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(id)
	}

	// Block main until done fires
	<-done
	fmt.Println("=== Stopped after 5 seconds ===")
}
