package gofaker

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	FIRSTNAMESFILE = "firstnames.txt"
	LASTNAMESFILE  = "lastnames.txt"
)

var (
	firstNames []string = []string{}
	lastNames  []string = []string{}
)

func loadData() error {

	file, err := os.Open(FIRSTNAMESFILE)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	fns := []string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fns = append(record)
	}
	firstNames = fns

	file, err = os.Open(LASTNAMESFILE)
	if err != nil {
		return err
	}
	defer file.Close()
	reader = csv.NewReader(file)
	lns := []string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		lns = append(record)
	}
	lastNames = lns

	return nil
}

func Name() string {
	switch rand.Intn(10) {
	case 0:
		return fmt.Sprintf("%v %v %v", Prefix(), FirstName(), LastName())
	case 1:
		return fmt.Sprintf("%v %v %v %v", Prefix(), FirstName(), LastName(), Suffix())
	case 2:
		return fmt.Sprintf("%v %v %v", FirstName(), LastName(), Suffix())
	default:
		return fmt.Sprintf("%v %v %v", Prefix(), FirstName(), LastName())
	}
	return ""
}

func FirstName() string {
	return firstNames[rand.Intn(len(firstNames))]
}

func LastName() string {
	return lastNames[rand.Intn(len(lastNames))]
}

func Prefix() string {
	p := []string{"Mr", "Mrs", "Ms", "Miss", "Dr"}
	return p[rand.Intn(len(p))]
}

func Suffix() string {
	s := []string{"Jr", "Sr", "I", "II", "III", "IV", "V", "MD", "DDS", "PhD", "DVM"}
	return s[rand.Intn(len(s))]
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	err := loadData()
	if err != nil {
		log.Println("gofaker: Error loading data files: ", err)
		os.Exit(1)
	}
}
