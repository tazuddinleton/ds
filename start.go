package main

import (
	"fmt"

	"github.com/tazuddinleton/ds/singly"
)

type Duck struct {
	Name string
}

func main() {
	// l := singly.NewList("Hello", "World", "!")
	// l2 := singly.NewList(1.0, 2.0, 3.0, 4.0, 4.09)

	l3 := singly.NewList(&Duck{"Donald"}, &Duck{"Trump"})
	fmt.Println(fmt.Printf("%v", singly.Flatten(l3)))
}
