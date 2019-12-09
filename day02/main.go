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

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			m := parse(b, i, j)
			if m == 19690720 {
				fmt.Println("\n Success")
				fmt.Printf("noun %02d , verb %02d", i, j)
				break
			}
		}
	}

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
	if op < 0 || op > 99 {
		return 0
	}

	if code < 0 || code > 99 {
		return 0
	}

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

	value := c[0]
	wipe(c)

	return value
}

func wipe(b []int) {
	for i := range b {
		b[i] = 0
	}
}
