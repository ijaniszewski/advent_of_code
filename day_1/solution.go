package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
        lines = append(lines, number)
    }
    return lines, scanner.Err()
}

func getTwoMultiplied(list_of_numbers []int) (int) {
    for _, first_number := range list_of_numbers {
        for _, second_number := range list_of_numbers {
            if first_number + second_number == 2020{
                return first_number * second_number
            }
		}
    }
    return 0
    }

func getThreeMultiplied(list_of_numbers []int) (int) {
    for _, first_no := range list_of_numbers {
        for _, second_no := range list_of_numbers {
            for _, third_no := range list_of_numbers {
                if first_no + second_no + third_no == 2020{
                    return first_no * second_no * third_no
                }
            }
		}
    }
    return 0
    }


func main() {
	var input_file string // variable declaration
	input_file = "input.txt"

    list_of_numbers, err := readLines(input_file)
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }
    var two_multiplied int
    two_multiplied = getTwoMultiplied(list_of_numbers)
    fmt.Println("Two multiplied: ", two_multiplied)

    var three_multiplied int
    three_multiplied = getThreeMultiplied(list_of_numbers)
    fmt.Println("Three multiplied: ", three_multiplied)


}