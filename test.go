package main

import (
	"log"
	"strconv"
)

type per struct {
	name string
}

func main() {
	m := make(map[string]*per)

	for i := 0; i < 5; i++ {
		if i % 2 == 0 {
			m[strconv.Itoa(i)] = &per{name:"yinuo"}
		} else {
			m[strconv.Itoa(i)] = nil
		}
	}

	for k, v := range m {
		if v != nil {
			log.Println(k, *v)
		}
	}
}