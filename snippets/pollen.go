package main

import (
	"math"
	"math/rand"
	"fmt"
	"time"
)

// Try to simulate brown Brownian motion

// Position stands up on itself
type Position struct {
	x, y float64
}

// move to a new position, distance 1, angle random
func (p *Position) move() {
	r := rand.Float32()
	xStep := rand.Float64()
	if r > .5 {
		p.x += xStep
	} else {
		p.x -= xStep
	}
	r = rand.Float32()
	yStep := math.Sqrt(1 - math.Pow(xStep, 2))
	if r > .5 {
		p.y -= yStep
	} else {
		p.y += yStep
	}
}

// Steps means the pollen moven n steps
func (p *Position) Steps(n int) float64 {
	if n < 0 {
		panic("n should be great than 0")
	}
	x0, y0 := p.x, p.y
	for i := 0; i < n; i++ {
		p.move()
	}
	return math.Sqrt(math.Pow((p.x - x0), 2) + math.Pow((p.y - y0), 2))
}

// Reset restore the zero value
func (p *Position) Reset() {
	p.x, p.y = 0, 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	p := Position{0, 0}
	fmt.Printf("N: %d, distance: %f\n", 1, p.Steps(1))
	p.Reset()
	fmt.Printf("N: %d, distance: %f\n", 2, p.Steps(2))
	p.Reset()
	fmt.Printf("N: %d, distance: %f\n", 10, p.Steps(10))
	p.Reset()
	fmt.Printf("N: %d, distance: %f\n", 100, p.Steps(100))
	p.Reset()
	fmt.Printf("N: %d, distance: %f\n", 1000, p.Steps(1000))
	p.Reset()
	fmt.Printf("N: %d, distance: %f\n", 10000, p.Steps(10000))
	p.Reset()
	fmt.Printf("N: %d, distance: %f\n", 100000, p.Steps(100000))
	p.Reset()
	fmt.Printf("N: %d, distance: %f\n", 1000000, p.Steps(1000000))
}