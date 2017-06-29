# Monty Hall (Golang) [![Build Status](https://travis-ci.org/edj-boston/monty-hall-go.svg?branch=master)](https://travis-ci.org/edj-boston/monty-hall-go) [![Coverage Status](https://coveralls.io/repos/github/edj-boston/monty-hall-go/badge.svg?branch=master)](https://coveralls.io/github/edj-boston/monty-hall-go?branch=master)
A Monty Hall problem simulator written in Golang

## Usage

```go
package main

import (
	"fmt"
	"github.com/edj-boston/monty-hall-go"
)

func main() {
	mh := montyhall.MontyHall{}
	mh.Run()

	fmt.Printf("%+v\n", mh.Results)
}
```

Will output:
```go
{Total:100000 Changes:50111 ChangeWins:33497 ChangeWinPercent:66.84560276186865 Stays:49889 StayWins:16725 StayWinPercent:33.524424221772335}
```
