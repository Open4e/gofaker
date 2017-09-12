package person

import (
	"fmt"
	"math/rand"
	"time"
)

type Prefixer interface {
	Prefix() string
}

type FirstNamer interface {
	FirstName() string
}

type LastNamer interface {
	LastName() string
}

type Suffixer interface {
	Suffix() string
}

type Namer interface {
	Prefixer
	FirstNamer
	LastNamer
	Suffixer
}

type Person struct {
	Data Data
}

func FullName(p Namer) string {
	switch rand.Intn(10) {
	case 0:
		return fmt.Sprintf("%v %v %v", p.Prefix(), p.FirstName(), p.LastName())
	case 1:
		return fmt.Sprintf("%v %v %v %v", p.Prefix(), p.FirstName(), p.LastName(), p.Suffix())
	case 2:
		return fmt.Sprintf("%v %v %v", p.FirstName(), p.LastName(), p.Suffix())
	default:
		return fmt.Sprintf("%v %v", p.FirstName(), p.LastName())
	}
}

func (p Person) FirstName() string {
	return p.Data.firstNames[rand.Intn(len(p.Data.firstNames))]
}

func (p Person) LastName() string {
	return p.Data.lastNames[rand.Intn(len(p.Data.lastNames))]
}

func (p Person) Prefix() string {
	return p.Data.prefixes[rand.Intn(len(p.Data.prefixes))]
}

func (p Person) Suffix() string {
	return p.Data.suffixes[rand.Intn(len(p.Data.suffixes))]
}

func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func init() {
	Seed()
}
