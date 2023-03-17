package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	if len(args) < 2 {
		return errors.New("expect second arg as type of getter")
	}
	fmt.Println(strings.Join(args, ", "))
	var s getter
	if args[1] == "with_race" {
		s = newServiceWithRaceAndRun(10)
	} else if args[1] == "without_race" {
		s = newServiceWithoutRaceAndRun(10)
	} else {
		return errors.New("expect second arg as `with_race` or `without_race`")
	}
	d := newDoer(s)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				if err := d.do(); err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	wg.Wait()
	return nil
}

type getter interface {
	get() string
}

func newDoer(g getter) *doer {
	return &doer{g: g}
}

type doer struct {
	g getter
}

func (d *doer) do() error {
	_ = d.g.get()
	n := time.Duration(rand.Int31n(200))
	time.Sleep(time.Microsecond * n)
	//fmt.Print(v)
	return nil
}
