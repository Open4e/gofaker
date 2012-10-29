package main

import (
	"fmt"
	"github.com/4eek/gofaker/lorem"
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

	fmt.Printf("\n%v\n", lorem.Word())
	fmt.Printf("\n%v\n", lorem.Words(1))
	fmt.Printf("\n%v\n", lorem.Words(5))
	fmt.Printf("\n%v\n", lorem.Sentence(10))
	fmt.Printf("\n%v\n", lorem.Sentences(10))
	fmt.Printf("\n%v\n", lorem.Paragraph(10))
	fmt.Printf("\n\n\n%v\n", lorem.Paragraphs(3))
}
