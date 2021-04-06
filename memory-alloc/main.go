package main

func main() {
	foo()
}

func foo() {
	println("a")
}

func main1() {
	_ = stackIt()
}

//go:noinline
func stackIt() int {
	y := 2
	return y * 2
}

func main2() {
	_ = stackIt2()
}

//go:noinline
func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
}

func main3() {
	y := 2
	_ = stackIt3(&y) // pass y down the stack as a pointer
}

//go:noinline
func stackIt3(y *int) int {
	res := *y * 2
	return res
}

func copyIt() {
	_ = CreateCopy()
}

func pointerIt() {
	_ = CreatePointer()
}

//go:noinline
func CreateCopy() BigStruct {
	return BigStruct{
		A: 123,
		B: 456,
		C: 789,
		D: "ABC",
		E: "DEF",
		F: "HIJ",
		G: true,
		H: true,
		I: true,
	}
}

//go:noinline
func CreatePointer() *BigStruct {
	return &BigStruct{
		A: 123,
		B: 456,
		C: 789,
		D: "ABC",
		E: "DEF",
		F: "HIJ",
		G: true,
		H: true,
		I: true,
	}
}

type BigStruct struct {
	A, B, C int
	D, E, F string
	G, H, I bool
}
