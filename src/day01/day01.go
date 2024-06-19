package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "strings"
)

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r))
	return string(r)
}

func readFile(filePath string) []int64 {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []int64
	for scanner.Scan() {
		depth, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, depth)
	}
	return lines
}

func GetIncreasedCount() int {
	lines := readFile("input.txt")
	return len(lines)
	// for i, v := range lines {

	// }
	// var count int
	// var prev rune
	// for _, r := range ReverseRunes("1234") {
	// 	if prev == 0 {
	// 		prev = r
	// 	} else if r < prev {
	// 		count++
	// 	}
	// 	prev = r
	// }
	// return count
}
