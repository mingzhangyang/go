package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

type Table chan *Ball

func Play(player string, table Table, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			ball := <-table
			ball.hits++
			fmt.Println(player, ball.hits)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		}
	}
}

func main() {
	table := make(Table)
	quit := make(chan bool, 2)
	go Play("Player 1 hits the ball, ", table, quit)
	go Play("Player 2 hits the ball, ", table, quit)

	table <- new(Ball) // game on, toss the ball
	time.Sleep(1 * time.Second)
	//<-table // game over, grab the ball
	go func() {
		quit <- true
		quit <- true
		// return
	}()
	go func() {
		close(table)
		// return
	}()
	panic("Show me the stacks...")
}
