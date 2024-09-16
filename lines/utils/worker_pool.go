package utils

import (
	"github.com/rs/zerolog/log"
	"sync"
)

func RunInWorkerPool(tasks chan func(), workersCount int) {
	var wg sync.WaitGroup
	wg.Add(workersCount)

	log.Printf("Running %v tasks in worker pool of size %v", len(tasks), workersCount)

	for i := 0; i < workersCount; i++ {
		go func() {
			defer wg.Done()

			for task := range tasks {
				task()
			}
		}()
	}

	wg.Wait()
	log.Printf("Finished running tasks in worker pool")
}

func RunInWorkerPoolReturn(tasks chan func() error, workersCount int) []error {
	var wg sync.WaitGroup
	wg.Add(workersCount)

	log.Printf("Running %v tasks in worker pool of size %v", len(tasks), workersCount)

	errors := make([]error, 0)

	for i := 0; i < workersCount; i++ {
		go func() {
			defer wg.Done()

			for task := range tasks {
				err := task()
				if err != nil {
					errors = append(errors, err)
				}
			}
		}()
	}

	wg.Wait()
	log.Printf("Finished running tasks in worker pool")

	return errors
}
