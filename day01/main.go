package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	data := readInput("input.txt")

	var total int

	for _, mass := range data {
		fuel := mass/3 - 2
		total += fuel

	}
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
