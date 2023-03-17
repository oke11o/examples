package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func newServiceAndRun(n int) *service {
	s := &service{n: n, d: make(map[int]*string, n)}
	for i := 0; i < n; i++ {
		s.d[i] = new(string)
	}
	s.update()
	go s.run()
	return s
}

type service struct {
	d map[int]*string
	n int
}

func (s *service) run() {
	t := time.NewTicker(time.Millisecond * 100)
	for range t.C {
		s.update()
	}
}

func (s *service) get() string {
	return *s.d[rand.Intn(s.n)]
}

func (s *service) update() {
	fmt.Println("update")
	for _, v := range s.d {
		newVal := strconv.Itoa(rand.Int())
		*v = newVal
	}
}
