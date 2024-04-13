package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	name            int
}

func (p Philo) eat(host chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 3; j++ {
		<-host
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("Philosopher %d is starting to eating\n", p.name)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Printf("Philosopher %d is finishing eating\n", p.name)
		host <- true
	}
}

func host(host chan bool) {
	for i := 0; i < 2; i++ {
		host <- true
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}

	hostCh := make(chan bool, 2)
	go host(hostCh)

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(hostCh, &wg)
	}
	wg.Wait()
	close(hostCh)
}
