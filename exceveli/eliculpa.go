func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        // Process the job (in this case, we'll just square the number)
        result := job * job // Processing the job
        results <- result   // Send the result to the results channel
    }
}
