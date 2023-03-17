package main

import (
	"math/rand"
	"strconv"
	"time"
)

func newServiceWithRaceAndRun(n int) *serviceWithRace {
	s := &serviceWithRace{n: n, d: make(map[int]*string, n)}
	for i := 0; i < n; i++ {
		s.d[i] = new(string)
	}
	s.update()
	go s.run()
	return s
}

type serviceWithRace struct {
	d map[int]*string
	n int
}

func (s *serviceWithRace) run() {
	t := time.NewTicker(time.Millisecond * 100)
	for range t.C {
		s.update()
	}
}

func (s *serviceWithRace) get() string {
	v := s.d[rand.Intn(s.n)]
	return *v
}

func (s *serviceWithRace) update() {
	for _, v := range s.d {
		newVal := strconv.Itoa(rand.Int())
		*v = newVal
	}
}
