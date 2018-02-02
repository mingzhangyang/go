package main

import (
	"fmt"
	"time"
)

type Job struct {
	i int
}

func (j Job) do() {
	<-time.After(1 * time.Second)
	fmt.Println(j.i)
}

// type Worker chan Job

// func (w Worker) run(q chan bool) {
// 	for {
// 		select {
// 		case j := <-w:
// 			j.do()
// 		case <-q:
// 			return
// 		}
// 	}
// }

// func createWorker(n int) Worker {
// 	c := make(chan Job, n)
// 	return c
// }

// type Pool struct {
// 	pool chan Worker
// 	n    int
// }

// func (p *Pool) Init(q chan bool) {
// 	p.pool = make(chan Worker, p.n)
// 	for i := 0; i < p.n; i++ {
// 		w := createWorker(10)
// 		p.pool <- w
// 		go w.run(q)
// 	}
// }

func process(c chan Job) {
	for j := range c {
		j.do()
	}
}

func main() {
	stop := make(chan bool)
	jobQueue := make(chan Job, 10)
	go func() {
		for i := 0; i < 100; i++ {
			jobQueue <- Job{i}
		}
		close(jobQueue)
		// close(stop) // if do not close stop, all goroutine will keep working for ever. Because closed jobQuequ always emit zero value.
	}()

	// fmt.Println(jobQueue)
	// wp := Pool{n: 10}
	// fmt.Println(wp)
	// wp.Init(stop)
	// fmt.Println(wp)

	// for {
	// 	j := <-jobQueue
	// 	w := <-wp.pool
	// 	w <- j
	// 	wp.pool <- w
	// }

	for i := 0; i < 10; i++ {
		go process(jobQueue)
	}

	<-time.After(time.Second * 15)
	close(stop)
}
