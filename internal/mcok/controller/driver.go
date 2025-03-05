package controller

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"
)

func CreateOutputDirIfNotExist(outputPath string) error {
	dirPath := filepath.Dir(outputPath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permissions
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteOutput(outputPath string, dataCount int64, generator func() string) error {
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

	maxConcurrent := 100
	var bufferSize int64 = min(100, dataCount)
	goroutineQueue := make(chan int, maxConcurrent)

	bar := progressbar.NewOptions(
		int(dataCount),
		progressbar.OptionSetWidth(40),
		progressbar.OptionShowTotalBytes(true),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n\n")
		}),
	)

	round := (dataCount / bufferSize) + 1
	fmt.Printf("Starting %d goroutines...\n\n", round)

	for i := range round {
		roundBuffer := bufferSize
		if i == round-1 {
			roundBuffer = dataCount % bufferSize
		}

		goroutineQueue <- 1
		bar.Add(int(roundBuffer))

		group.Go(func() error {
			// Consume from queue for other to continue when done
			defer func() {
				<-goroutineQueue
			}()

			var buffer bytes.Buffer

			for range roundBuffer {
				fmt.Fprintln(&buffer, generator())
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
