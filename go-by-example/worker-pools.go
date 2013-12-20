package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    jobs    := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        // 3 workers
        // passing in the current value of w as the id along with
        // the jobs and results channels.
        go worker(w, jobs, results)
    }

    // create 9 jobs
    for j := 1; j <= 9; j++ {
        // add a job to the jobs channel which is picked up by a worker
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 9; a++ {
        <-results
    }

}
