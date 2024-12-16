package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {

	t := time.Now() // Record the start time

	links := []string{
		"http://google.com",
		"http://linkedin.com",
		"http://nimendra.xyz",
		"http://github.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://example.com",
		"https://tailwindcss.com/docs/installation",
	}

	// Create a tabwriter for formatting output
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)

	// Print the table header
	fmt.Fprintln(w, "Link\tStatus\tResponse Time")

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1) // Increment the WaitGroup counter
		go func(link string) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			checkLink(w, link)
		}(link)

		// checkLink(w, link)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Flush the writer to ensure all output is displayed
	w.Flush()

	fmt.Printf("Total Time elapsed: %v\n", time.Since(t))
}

func checkLink(w *tabwriter.Writer, link string) {
	start := time.Now() // Start timing
	res, err := http.Get(link)
	elapsed := time.Since(start) // Calculate elapsed time

	if err != nil {
		fmt.Fprintf(w, "%s\tError: %s\n", link, err)
		return
	}
	defer res.Body.Close()

	fmt.Fprintf(w, "%s\t%s\t%v\n", link, res.Status, elapsed)
}
