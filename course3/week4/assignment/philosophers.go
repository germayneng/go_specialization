/*
The trick is to make use of buffer channel to limit the number of people starting to eat to avoid deadlocks! For e.g
if guy 1 starts eating, he will perform lock. At this point, we should push a message down the channel. This will take
up 1 buffered slot from the channel. Next concurrency will also try to lock. Neighbour of guy 1 will be blocked so
the only possible guy to eat are not neighbours of 1. Then when the guy starts eating, and lock, it will send another
message down the channel. Now all buffer channel is used up and so any sending is blocked, Remember when we eat, we need
to send to a channel so no one will be able to eat.

Once guy 1 and another guy are done - finished, and unlock, they should receive information from channel to unlock the
slot so other guy can start eating + lock + send to channel
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// create chopsticker struct
type chopS struct {
	sync.Mutex
}

// create philosopher struct
type Philo struct {
	leftCS *chopS
	rightCS *chopS
	number int
}

// method for philo
func (p *Philo) eat(ch chan int, wg *sync.WaitGroup) {
	// 3 times instead of infinite
	for i := 0; i<3; i++{
		// push number down channel to block a slot
		ch <- p.number

		// random lock
		// 50 % chance
		shuffle := rand.Intn(2)
		if shuffle == 1 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		}else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}

		fmt.Printf("\nPhilosoper %d is  staring to eat!  \n",p.number)
		fmt.Printf("Philosoper %d finished eating! \n",p.number)

		// unlock based on how we lock
		if shuffle == 1 {
			p.rightCS.Unlock()
			p.leftCS.Unlock()
		} else {
			p.leftCS.Unlock()
			p.rightCS.Unlock()
		}

		//once phil done eating 1 round, receive from ch to unblock a slot
		<- ch

	}




	wg.Done()

}



// we will use buffered channel to ensure only 2 philosopher can eat
// only 2 concurrency
func main() {

	// initialize chopsticks (aka number of locks)
	// create slice to store chopsticks
	Cstcks := make([]*chopS,5)

	for i := 0;i < 5; i++ {
		Cstcks[i] = &chopS{}
	}

	//create channel with buffer size 2
	ch := make(chan int,2)
	wg := sync.WaitGroup{}

	wg.Add(5) // 5 concurrency/workers must decrease it back

	//create 5 philosopers
	philos := make([]*Philo,5)
	for i := 0;i < 5; i++ {
		philos[i] = &Philo{
			leftCS: Cstcks[i],
			rightCS:Cstcks[(i+1)%5],
			number: i,
		}

		go philos[i].eat(ch,&wg)
	}


	// block until all cleared
	wg.Wait()

}
