package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count = 0

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Simulate a compute-heavy operation
	go func() {
		// Simulate a compute-heavy operation by sleeping for a duration
		count++
		if count%5 == 0 {
			fmt.Println("Sleeping")
			time.Sleep(5 * time.Second)
		}

		fmt.Println("Compute operation complete!")
	}()

	// Respond immediately to the client
	fmt.Fprint(w, "Request received and processing...\n")
}
