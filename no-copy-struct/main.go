package main

import (
	"fmt"
	"sync"
)

type EmbedMutex struct {
	sync.Mutex
	ID int
}

type noCopy struct{}

func (i *noCopy) Lock()   {}
func (i *noCopy) Unlock() {}

type EmbedNoCopy struct {
	noCopy
	ID int
}

func main() {
	// m := EmbedMutex{ID: 1}
	// m1 := m
	// fmt.Println(m, m1)

	nc := EmbedNoCopy{ID: 2}
	nc1 := nc
	fmt.Println(nc, nc1)
}

/*
$ go vet no-copy-struct/main.go
# command-line-arguments
no-copy-struct/main.go:29:9: assignment copies lock value to nc1: command-line-arguments.EmbedNoCopy
no-copy-struct/main.go:30:14: call of fmt.Println copies lock value: command-line-arguments.EmbedNoCopy
no-copy-struct/main.go:30:18: call of fmt.Println copies lock value: command-line-arguments.EmbedNoCopy
*/
