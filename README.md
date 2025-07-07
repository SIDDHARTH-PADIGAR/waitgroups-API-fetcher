## Use Case: Coordinating Multiple Concurrent Web Fetches

You're scraping or pinging multiple endpoints, and you want to do it fast. You don’t want to fetch each URL one by one, instead, you fire off all requests concurrently and wait until **all of them finish**, before moving on.

---

## What's Going On:

* A `WaitGroup` is initialized to track when all goroutines complete.
* For each URL, the counter is incremented using `wg.Add(1)`.
* A separate goroutine is launched for each fetch using `go fetchURL(...)`.
* Inside each goroutine:

  * It performs an HTTP GET request.
  * It defers `wg.Done()` to signal task completion.
* The main function calls `wg.Wait()` which blocks until **every single fetch has completed**.

---

## Why It Worked

* Each fetch runs in a separate goroutine, which allows all network calls to happen concurrently.
* The `WaitGroup` acts like a sync barrier, the main goroutine halts at `wg.Wait()` until every worker calls `wg.Done()`.
* This ensures we don’t exit the program early while other goroutines are still doing their job.

> Without `Wait()`, the main function would exit before any or all of the fetches finish and you'd see nothing, or partial output at best.

### Output

```
Fetched: https://example.com Status Code: 200
Fetched: https://golang.org Status Code: 200
Fetched: https://httpbin.org/get Status Code: 200
All fetches complete.
```

---

## When To Use This Pattern

### Coordinating Concurrent Tasks

* Running batch jobs in parallel (e.g., hitting multiple APIs)
* Performing parallel computations (e.g., image processing, file I/O)
* Waiting for worker goroutines to finish before shutdown or next step

### Preventing Premature Exit

* Useful in scripts and CLI tools where the main thread might exit before background tasks are done
* Ensures clean and complete execution in concurrent systems
