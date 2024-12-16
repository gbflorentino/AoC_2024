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

// Part 1
func verifyReports(reports [][]int) int {
	var safeReports int = 0
	
	for _, report := range reports {
		var reportSafeness int = 0
		var lastLevel, lastDiference int = 0, 0

		for _, level := range report {
			diference := level - lastLevel

			if (lastLevel == 0) {
				lastLevel = level
				continue
			} else if (diference * lastDiference < 0) {
				reportSafeness = 0
				break
			} else if (diference == 0) {
				reportSafeness = 0
				break				
			} else if (intAbs(diference) < 1 || intAbs(diference) > 3) {
				reportSafeness = 0
				break
			}

			lastLevel, lastDiference = level, diference
			reportSafeness = 1
		}

		if (reportSafeness != 0) {
			safeReports += 1
		}
	}

	return safeReports
} 

// Part 2
func verifyReport(report []int) (int, []int) {
	var unsafeLevels []int
	var reportSafeness int = 0
	var lastLevel, lastDiference int = 0, 0

	for index, level := range report {
		diference := level - lastLevel

		if (lastLevel == 0) {
			lastLevel = level
			continue
		} else if (diference * lastDiference < 0) {
			unsafeLevels = append(unsafeLevels, index-1)
			reportSafeness += 1
		} else if (diference == 0) {
			unsafeLevels = append(unsafeLevels, index)
			reportSafeness += 1
		} else if (intAbs(diference) < 1 || intAbs(diference) > 3) {
			unsafeLevels = append(unsafeLevels, index)
			reportSafeness += 1
		}

		lastLevel, lastDiference = level, diference
	}

	return reportSafeness, unsafeLevels
}

func verifyReportsDampener(reports [][]int) int {
	var safeReports int = 0

	for _, report := range reports {

		fmt.Print(report, " ")

		reportSafeness, unsafeLevels := verifyReport(report)
		
		if (reportSafeness == 0) {
			safeReports += 1
		} else if (reportSafeness <= 2) {
			removeIndex := unsafeLevels[0]
			report = append(reportSafeness[:removeIndex], reportSafeness[removeIndex+1:]...)
			reportSafeness, _ = verifyReport(report)
			
		}

		fmt.Println(reportSafeness, unsafeLevels)
	}

	return safeReports
} 

func main(){
	filepath := "inputs/day_2_p1_test"

	reports := readInput(filepath)
	i := verifyReportsDampener(reports)
	fmt.Println(i)
}
