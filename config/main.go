package main

import (
	"config/env"
	"fmt"
)

func main() {
	fmt.Printf("ReadConfigViper:\t%+v\n", env.ReadConfigViper())
	fmt.Printf("ReadConfigViperMarshal:\t%+v\n", env.ReadConfigViperMarshal())
}
