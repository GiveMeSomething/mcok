package mcok

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Idea:
// Prepare a buffer to pre-construct a small part of the result
// When done, try to acquire the lock and write to file

func writeOutput(outputPath string, round int64) error {
	lock := sync.Mutex{}

	group, _ := errgroup.WithContext(context.Background())

	outputFile, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	err = os.Truncate(outputPath, 0)
	if err != nil {
		return err
	}

	maxConcurrent := 200
	var bufferSize int64 = 100
	goroutineQueue := make(chan int, maxConcurrent)

	fmt.Printf("Starting %d goroutines...\n", max(round/bufferSize, 1))

	for range max(round/bufferSize, 1) {
		goroutineQueue <- 1

		group.Go(func() error {
			defer func() {
				<-goroutineQueue
			}()

			var buffer bytes.Buffer

			for range bufferSize {
				fmt.Fprintln(&buffer, randomPassenger().ToString())
			}

			lock.Lock()
			defer lock.Unlock()

			// Write from temp buffer to output
			_, err = outputFile.Write(buffer.Bytes())
			if err != nil {
				return err
			}

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
