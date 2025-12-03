package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const results = true
const verbose = true

func main() {
	st := time.Now()
	d2()

	fmt.Printf("\n%.2fs elapsed\n", time.Since(st).Seconds())
}

func readFileAsLines(day, part int) []string {
	f, err := os.OpenFile(fmt.Sprintf("inputs/%v_%v.txt", day, part), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	con, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(con), "\r\n")
}

type level int

const (
	d level = iota
	r
)

func p(l level, vals ...any) {
	if verbose && l == d {
		fmt.Println("DEBUG:", fmt.Sprint(vals...))
	} else if results && l == r {
		fmt.Println("RESULT:", fmt.Sprint(vals...))
	}
}
