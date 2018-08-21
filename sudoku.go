package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//config := Read_config("test_config_file.txt")

}

func PrintSquare() (square string) {
	empty_line :=
		" _  _  _ | _  _  _ | _  _  _\n"
	separator_line := "---------|---------|---------\n"
	square = "\n"
	square += get_square(empty_line, separator_line)

	return
}

func get_square(empty_line string, separator_line string) (square string) {
	for triple_row_idx := 0; triple_row_idx < 3; triple_row_idx++ {
		square += get_row(empty_line)
		if triple_row_idx == 2 {
			continue
		}
		square += separator_line
	}
	return
}

func get_row(empty_line string) (square string) {
	for empty_line_idx := 0; empty_line_idx < 3; empty_line_idx++ {
		square += empty_line
	}
	return square
}

func Read_config(path string) (content string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
		content += "\n"
	}
	fmt.Println(content)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
