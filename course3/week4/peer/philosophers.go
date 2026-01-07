package main

import (
	"fmt"
	"sync"
)

func main() {
	host := &Host{
		turns: make(chan bool, 2),
	}
	host.getReady()

	sticks := make([]*ChopStick, 5)
	for i := 0; i <  5; i++ {
		sticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i <  5; i++ {
		philosophers[i] = &Philosopher{
			left: sticks[i],
			right: sticks[(i + 1) % 5],
			name: i + 1,
			eatenTimes: 0,
			host: host,
		}
	}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(5)
	for _, philosopher := range philosophers {
		philosopher := philosopher
		go func() {
			philosopher.startDinning()
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

type ChopStick struct{ sync.Mutex }

type Philosopher struct{
	left, right *ChopStick
	name, eatenTimes int
	host *Host
}

type Host struct {
	turns              chan bool
}

func (h Host) canPhilosopherEat() bool {
	return <- h.turns
}

func (h Host) philosopherFinishedEating() {
	h.turns <- true
}

func (h Host) getReady() {
	h.turns <- true
	h.turns <- true
}

func (p *Philosopher) startDinning(){
	for p.isHungry(){
		if p.host.canPhilosopherEat() {
			p.eat()
			p.host.philosopherFinishedEating()
		}
	}
}

func (p *Philosopher) eat()  {
	p.left.Lock()
	p.right.Lock()
	fmt.Printf("starting to eat %d\n", p.name)
	p.eatenTimes++
	fmt.Printf("finishing eating %d\n", p.name)
	p.left.Unlock()
	p.right.Unlock()
}

func (p *Philosopher) isHungry() bool{
	return p.eatenTimes < 3
}
