package gtools

import (
	"sync"
	"testing"
	"time"
)

func TestRandomTool(t *testing.T) {
	random1 := NewRandom()
	time.Sleep(1 * time.Millisecond)
	random2 := NewRandom()

	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Log(
				random1.GenString(6), random2.GenString(6),
				random1.GenCode(6), random2.GenCode(6),
				random1.GenNum(10, 100), random2.GenNum(10, 100),
			)
		}()
	}
	wg.Wait()
}
