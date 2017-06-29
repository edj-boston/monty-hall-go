package main

import (
	"fmt"
	"github.com/edj-boston/monty-hall-go/lib/MontyHall"
)

func main() {
	mh := montyhall.MontyHall{}
	mh.Run()

	fmt.Printf("%+v\n", mh.Results)
}
