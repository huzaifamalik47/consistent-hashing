package main

import (
	"fmt"
	"reflect"
)

type Ring struct {
	size int
}

func main() {

	var servers = make(map[string]int, 0)

	fmt.Println(reflect.TypeOf(servers))

	ring := NewRing(10)
	fmt.Println(ring)

}

func NewRing(size int) *Ring {
	return &Ring{
		size: size,
	}
}
