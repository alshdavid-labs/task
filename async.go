package task

import (
	"fmt"
	"sync"
)

func Async(ts ...*Task) bool {
	var wg sync.WaitGroup
	wg.Add(len(ts))

	ok := true
	for i := range ts {
		go func(t *Task, wg *sync.WaitGroup) {
			err := t.Do()
			if err != nil {
				fmt.Println(err)
				ok = false
			}
			wg.Done()
		}(ts[i], &wg)
	}

	wg.Wait()
	return ok
}

func AsyncSilent(ts ...*Task) {
	var wg sync.WaitGroup
	wg.Add(len(ts))

	for i := range ts {
		go func(t *Task, wg *sync.WaitGroup) {
			t.Do()
			wg.Done()
		}(ts[i], &wg)
	}

	wg.Wait()
}
