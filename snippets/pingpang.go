package main

import (
	"fmt"
	"time"
)

type ball struct {
	hits int
}

type table chan *ball

func play(player string, tab table, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			bal := <-tab
			bal.hits++
			fmt.Println(player, bal.hits)
			time.Sleep(100 * time.Millisecond)
			tab <- bal
		}
	}
}

func main() {
	defer func() {
		fmt.Println("Done")
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	tab := make(table)
	quit := make(chan bool, 2)
	go play("Player 1 hits the ball, ", tab, quit)
	go play("Player 2 hits the ball, ", tab, quit)

	tab <- new(ball) // game on, toss the ball
	time.Sleep(1 * time.Second)
	//<-table // game over, grab the ball
	go func() {
		quit <- true
		quit <- true
		// return
	}()
	go func() {
		close(tab)
		// return
	}()
	panic("Show me the stacks...")
}
