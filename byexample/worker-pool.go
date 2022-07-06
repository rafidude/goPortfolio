package main

import (
	"fmt"
	"time"
)

// worker receives jobs and sends results back to the pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// receive results from the workers
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
