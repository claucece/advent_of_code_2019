package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b := readFileToString("input.txt")
	n := parse(b, 12, 2)
	fmt.Printf("%d \n", n)
}

func readFileToString(fname string) []int {
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.TrimSpace(string(bytes))
	ss := strings.Split(s, ",")

	var bb []int
	for _, b := range ss {
		sb, err := strconv.Atoi(b)
		if err != nil {
			return nil
		}
		bb = append(bb, sb)
	}

	return bb
}

func parse(bytes []int, op int, code int) int {
	c := make([]int, len(bytes))
	copy(c[:], bytes[:])

	c[1] = op
	c[2] = code

	for i := 0; i < len(c); i += 4 {
		if c[i] == 99 {
			break
		} else if c[i] == 1 {
			x, y, pos := c[i+1], c[i+2], c[i+3]
			add := c[x] + c[y]
			c[pos] = add

			continue
		} else if c[i] == 2 {
			x, y, pos := c[i+1], c[i+2], c[i+3]
			mul := c[x] * c[y]
			c[pos] = mul

			continue
		} else {
			break
		}
	}

	return c[0]
}
