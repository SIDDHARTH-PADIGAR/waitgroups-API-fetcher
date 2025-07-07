package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchURL(url string, wg *sync.WaitGroup) { //We pass a pointer so all goroutines share the same WaitGroup instance
	defer wg.Done()

	resp, err := http.Get(url) //Makes the HTTP request
	if err != nil {            //Checks for error, if there's one, logs and returns
		return
	}
	defer resp.Body.Close() //Makes sure the response body is closed after the function ends(cleanup)

	fmt.Println("Fetched:", url, "Status Code:", resp.StatusCode) //If successfu, prints the status code
}

func main() {
	var wg sync.WaitGroup //Initializes a new WaitGroup instance

	urls := []string{ //Declare a slice of URLs to fetch
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
	}

	for _, url := range urls { //Loop through each URL
		wg.Add(1)             //Tells the WaitGroup there will be one more goroutine
		go fetchURL(url, &wg) //Launch the fetch in a separate goroutine cause we want them to run concurrently, not one after another
	}

	wg.Wait()                            // This blocks the main function until all goroutines call wg.Done()
	fmt.Println("All fetches completed") //After all fetched are done, print that everything's complete
}
