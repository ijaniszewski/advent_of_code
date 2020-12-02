package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strings"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
        lines = append(lines, line)
	}
	
    return lines, scanner.Err()
}

func checkPassword(line string) (int, int) {
	var counter_1 int = 0
	var counter_2 int = 0
	s := strings.Split(line, " ")
	minmax := strings.Split(s[0], "-")
	var min, _ = strconv.Atoi(minmax[0])
	var max, _ = strconv.Atoi(minmax[1])
	var letter = strings.Trim(s[1], ":")
	var password = string(s[2])
	var letter_counted = strings.Count(password, letter)
	if min <= letter_counted {
		if letter_counted <= max {
			counter_1 = 1
		}
	}
	if string(password[min-1]) == letter {
		if string(password[max-1]) == letter {

		} else {
			counter_2 = 1
		}
	} 
	if string(password[min-1]) != letter {
		if string(password[max-1]) == letter {
			counter_2 = 1
		}
	}	
	
	return counter_1, counter_2
}




func main() {
	var input_file string // variable declaration
	input_file = "input.txt"

    lines, err := readLines(input_file)
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }

	var first_counter int
	var second_counter int
	first_counter = 0
	second_counter = 0
	for _, line := range lines {
		var cnt_1 int
		var cnt_2 int
		cnt_1, cnt_2 = checkPassword(line)
		first_counter += cnt_1
		second_counter += cnt_2
	}
	fmt.Println(first_counter, second_counter)

}