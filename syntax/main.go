package main

import (
	"fmt"
	"os"

	"github.com/fukuiretu/go-practice/syntax/shuffle"
)

type Dice []int

func (d Dice) Seed() int64   { return int64(os.Getpid()) }
func (d Dice) Len() int      { return len(d) }
func (d Dice) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

var dice = Dice([]int{1, 2, 3, 4, 5, 6})

type foo struct {
	v int
}

func (f *foo) add(v int) {
	f.v += v
}

func main() {
	fmt.Printf("%v\n", dice)
	shuffle.Shuffle(dice)
	fmt.Printf("%v\n", dice)

	// NOTE: ポインタの例
	a := &foo{v: 2}
	a.add(3)
	println("a.v:", a.v)
	a.add(4)
	println("a.v:", a.v)

	// NOTE: &foo{}と式new(foo)は等価
	a = new(foo)
	a.add(9)
	println("a.v:", a.v)
	a.add(1)
	println("a.v:", a.v)
}
