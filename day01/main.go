package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data := readInput("input.txt")

	var total int
	var gTotal int

	for _, mass := range data {
		total += mass/3 - 2

		for {
			fuel := mass/3 - 2
			if fuel <= 0 {
				break
			}
			gTotal += fuel
			mass = fuel
		}
	}

	fmt.Printf("\n %d \n", total)
	fmt.Printf("\n %d \n", gTotal)
}

func readInput(fname string) []int {
	file, err := os.Open("input.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var data []int
	for scanner.Scan() {
		result, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil
		}

		data = append(data, result)
	}

	return data
}
