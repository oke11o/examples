package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func newServiceWithoutRaceAndRun(n int) *serviceWithoutRace {
	s := &serviceWithoutRace{n: n, d: make(map[int]string, n)}
	s.update()
	go s.run()
	return s
}

type serviceWithoutRace struct {
	d map[int]string
	n int
	m sync.RWMutex
}

func (s *serviceWithoutRace) run() {
	t := time.NewTicker(time.Millisecond * 100)
	for range t.C {
		s.update()
	}
}

func (s *serviceWithoutRace) get() string {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.d[rand.Intn(s.n)]
}

func (s *serviceWithoutRace) update() {
	fmt.Println("update")
	d := make(map[int]string, len(s.d))
	for k := range d {
		newVal := strconv.Itoa(rand.Int())
		d[k] = newVal
	}
	s.m.Lock()
	s.d = d
	s.m.Unlock()
}
