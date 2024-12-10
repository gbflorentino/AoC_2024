package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func intAbs(num int) int {
	if num < 0 {
		return num * (-1)
	}
	return num
}

func listAtoi(stringSlice []string) []int {
	var intSlice []int

	for _, num := range stringSlice {
		num, _ := strconv.Atoi(num)
		intSlice = append(intSlice, num)
	}

	return intSlice
}

func readInput(filepath string) [][]int {
	var reports [][]int
	
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splittedLines := strings.Fields(line)
		reports = append(reports, listAtoi(splittedLines))
	}

	return reports
}

func verifyReports(reports [][]int) int {
	var safeReports int = 0
	
	for _, report := range reports {
		lastLevel, lastDiference := 0, 0
		for _, level := range report {
			diference := absInt(level - lastLevel)
			}			
		}
	}

	return 0
}

func main(){
	filepath := "inputs/day_2_p1_test"

	reports := readInput(filepath)
	i := verifyReports(reports)
	i+=1
}
