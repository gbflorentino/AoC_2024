package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

func absInt(num int) int {
	if num < 0 {
		num = num * (-1)
	}
	return num
}

func readColFile(filepath string) ([]int, []int) {
	var first_column, second_column []int
	
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Fields are used to split strings around one, or more, consecutive spaces
		splited_line := strings.Fields(line)

		first_num, _ := strconv.Atoi(splited_line[0])
		second_num, _ := strconv.Atoi(splited_line[1])

		first_column = append(first_column, first_num)
		second_column = append(second_column, second_num)
	}

	file.Close()

	// Sort the lists
	sort.Ints(first_column)
	sort.Ints(second_column)

	return first_column, second_column
}

// Part 1
func computeDistances(first_column []int, second_column []int) int {
	var distances int = 0
	
	for i := 0; i < len(first_column); i++ {
		distance := absInt(first_column[i] - second_column[i])
		distances = distances + distance
	}

	return distances
}

// Part 2
func computeSimilarities(first_column []int, second_column []int) int {
	var similarities int = 0
	
	for i := 0; i < len(first_column); i++ {
		repetition_score := 0
		for j := 0; j < len(second_column); j++ {
			if second_column[j] == first_column[i] {
				repetition_score = repetition_score + 1
			}
		}
		
		similarty_score := (first_column[i] * repetition_score)
		similarities = similarities + similarty_score
	}
	
	return similarities
}

func main() {
	filepath := "inputs/day_1_p1.txt"

	first_column, second_column := readColFile(filepath)

	distances_sum := computeDistances(first_column, second_column)
	fmt.Println("Part 1 Anwser:", distances_sum)

	similarity_sum := computeSimilarities(first_column, second_column)
	fmt.Println("Part 1 Anwser:", similarity_sum)
}
