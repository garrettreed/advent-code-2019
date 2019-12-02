package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (d Days) Day01() {
	file, err := os.Open("inputs/day01.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	var rockets []float64
	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 64)
		check(err)
		rockets = append(rockets, mass)
	}

	fmt.Println("Answer 1:", getSingleFuel(rockets))
	fmt.Println("Answer 2:", getTotalFuel(rockets))
}

func getSingleFuel(rockets []float64) int {
	var total float64
	for _, mass := range rockets {
		total += math.Floor(mass / 3) - 2
	}
	return int(total)
}

func getTotalFuel(rockets []float64) int {
	var total float64
	for _, mass := range rockets {
		for mass > 0 {
			mass = math.Floor(mass / 3) - 2
			if mass > 0 {
				total += mass
			}
		}
	}
	return int(total)
}