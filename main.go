package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

var eatTimes = 3
var eatDelay = 5 * time.Second

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{"Kant", 0, 1},
	{"Plato", 1, 2},
	{"Hume", 2, 3},
	{"Kierkegaard", 3, 4},
	{"Nietzsche", 4, 0},
}

func diningPrblm(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex) {
	defer wg.Done()

	for i := 0; i < eatTimes; i++ {
		if p.leftFork > p.rightFork {
			forks[p.rightFork].Lock()
			fmt.Printf("%s takes the right fork\n", p.name)
			forks[p.leftFork].Lock()
			fmt.Printf("%s takes the left fork\n", p.name)
		} else {
			forks[p.leftFork].Lock()
			fmt.Printf("%s takes the left fork\n", p.name)
			forks[p.rightFork].Lock()
			fmt.Printf("%s takes the right fork\n", p.name)
		}

		fmt.Printf("%s eating\n", p.name)
		time.Sleep(eatDelay)

		forks[p.rightFork].Unlock()
		forks[p.leftFork].Unlock()
	}
}

func dine() {
	wg.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningPrblm(philosophers[i], &wg, forks)
	}

	wg.Wait()
}

func main() {
	fmt.Println("Dining Philosophers!")

	dine()
}
