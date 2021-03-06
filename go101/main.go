package main

import (
	"fmt"
	"golds/oops"
)

func main() {
	fmt.Println("Hello GoLang")
	beyondHello()

	creature := oops.Creature{"Anuj", true}
	creature.Dump()
	runner := oops.Runner(creature)

	flyingCreature := oops.FlyingCreature{creature, false}
	flyingCreature.Dump()

	runner.Run()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)

}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}
