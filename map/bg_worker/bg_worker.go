package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	err := run(os.Args) // , os.Stdin, os.Stdout, os.Stderr
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	fmt.Println("DONE")
}

func run(args []string) error {
	s := newServiceAndRun(10)
	d := newDoer()
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				if err := d.do(s.get()); err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	wg.Wait()
	return nil
}

func newDoer() *doer {
	return &doer{}
}

type doer struct {
}

func (g *doer) do(v string) error {
	n := time.Duration(rand.Int31n(200))
	time.Sleep(time.Microsecond * n)
	//fmt.Print(v)
	return nil
}
