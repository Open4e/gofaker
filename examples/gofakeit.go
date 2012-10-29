package main

import (
	"fmt"
	"github.com/4eek/gofaker/person"
	"time"
)

func main() {
	num := 10000
	t0 := time.Now()
	for i := 0; i < num; i++ {
		fmt.Printf("%v\n", person.FullName())
	}
	t1 := time.Now()
	fmt.Printf("\n%v Usernames generated and printed to stdout in %v (%v)\n", num, t1.Sub(t0), t1.Sub(t0).Seconds())

	num = 10000
	t0 = time.Now()
	for i := 0; i < num; i++ {
		person.FullName()
	}
	t1 = time.Now()
	fmt.Printf("\n%v Usernames generated in %v (%v)\n", num, t1.Sub(t0), t1.Sub(t0).Seconds())
}
