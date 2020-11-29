package main

import "fmt"

type person struct {
	name string
}

//Printable ...
type Printable interface {
	print()
}

type figure struct {
	name string
}

type close func()

func invokeClose(c close) {
	c()
}

func (f *figure) print() {
	fmt.Println("[figure]", f.name)
}

func (p *person) print() {
	fmt.Println("[person]", p.name)
}

func invokePrint(p Printable) {
	p.print()
}

func (p *person) clean() {
	p.name = ""
}

func main() {
	p := &person{name: "Juan"}
	invokePrint(p)

	f := &figure{name: "Cubo"}
	invokePrint(f)

	function := func() {
		fmt.Println("Hello World")
	}
	invokeClose(function)

}
