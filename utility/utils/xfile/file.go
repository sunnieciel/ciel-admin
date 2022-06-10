package xfile

import (
	"bufio"
	"os"
)

// ReadLine read line from file
func ReadLine(path string, lineNum int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNum {
			return fileScanner.Text(), nil
		}
		lineCount++
	}
	if err = file.Close(); err != nil {
		return "", err
	}
	return "", nil
}
