package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 10

func main() {
	// Initialize the shared buffer
	buffer := make([]byte, bufferSize)

	// Channels for signaling read and write operations
	readCh := make(chan struct{})
	writeCh := make(chan struct{})

	// Mutex for protecting the shared buffer
	var mutex sync.Mutex

	// Number of reading and writing goroutines
	M := 8
	N := 2

	// Start writing goroutines
	for i := 0; i < N; i++ {
		go writer(i, buffer, writeCh, &mutex)
	}

	// Start reading goroutines
	for i := 0; i < M; i++ {
		go reader(i, buffer, readCh, writeCh, &mutex)
	}

	// Keep the main goroutine running
	select {}
}

func writer(id int, buffer []byte, writeCh chan struct{}, mutex *sync.Mutex) {
	for {
		// Acquire the lock before writing
		mutex.Lock()

		// Simulate writing to the buffer
		fmt.Printf("Writer %d writing\n", id)
		time.Sleep(time.Millisecond * 500)

		// Release the lock after writing
		mutex.Unlock()

		// Signal that a write operation has occurred
		writeCh <- struct{}{}
	}
}

func reader(id int, buffer []byte, readCh, writeCh chan struct{}, mutex *sync.Mutex) {
	for {
		// Wait for a signal that a write has occurred
		<-writeCh

		// Acquire the lock before reading
		mutex.Lock()

		// Simulate reading from the buffer
		fmt.Printf("Reader %d reading\n", id)
		time.Sleep(time.Millisecond * 200)

		// Release the lock after reading
		mutex.Unlock()

		// Signal that a read operation has occurred
		readCh <- struct{}{}
	}
}
