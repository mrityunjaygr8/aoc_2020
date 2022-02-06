package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Read_file(filename string) []string {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Println("Error reading input file")
	}

	b := bufio.NewScanner(f)
	b.Split(bufio.ScanLines)

	var fileTextLines []string

	for b.Scan() {
		fileTextLines = append(fileTextLines, b.Text())
	}

	return fileTextLines
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func Atoi(s string) (int, error) {
	tmp, err := strconv.Atoi(s)
	return tmp, err
}
