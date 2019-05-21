package main

import (
	"fmt"
)
func Thirty() {
	jobs := make(chan int, 5)
	done := make(chan bool)


	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("Received Job ", j)
			} else {
				fmt.Println("Received all jobs") 
				done <- true
				return 
			}
		}
	}()
	for j:= 1; j <= 4; j++ {
		jobs <- j
		fmt.Println("Job Sent ", j)
	}
	close(jobs)
	fmt.Println("all jobs Sent")

	<-done
}