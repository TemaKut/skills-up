package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ms := GetMetricsService()
			println(ms.Id())
		}()
	}

	wg.Wait()
}
