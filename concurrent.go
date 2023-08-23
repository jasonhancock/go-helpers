package helpers

import "sync"

// ConcurrentWork will execute fn for every input in the work slice, doing a
// maximum of maxConcurrency concurrently.
func ConcurrentWork[T any](work []T, fn func(T), maxConcurrency int) {
	ch := make(chan T, len(work))
	var wg sync.WaitGroup

	for i := 0; i < maxConcurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range ch {
				fn(msg)
			}
		}()
	}

	for _, msg := range work {
		ch <- msg
	}

	close(ch)
	wg.Wait()
}
