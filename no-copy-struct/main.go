package main

import "fmt"

type T struct {
	noCopy
	A, B int
}

func print(t T) {
	fmt.Printf("%d%d", t.A, t.B)

}
func main() {
	a := T{A: 1, B: 9}
	b := a
	print(a)
	b.A = 2
	b.B = 6
	print(b)
}

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

/*
$ go run no-copy-struct/main.go
1926

$ go vet no-copy-struct/main.go
# command-line-arguments
no-copy-struct/main.go:29:9: assignment copies lock value to nc1: command-line-arguments.EmbedNoCopy
no-copy-struct/main.go:30:14: call of fmt.Println copies lock value: command-line-arguments.EmbedNoCopy
no-copy-struct/main.go:30:18: call of fmt.Println copies lock value: command-line-arguments.EmbedNoCopy
*/
