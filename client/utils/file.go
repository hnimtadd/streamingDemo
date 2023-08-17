package utils

import (
	"bufio"
	"os"
)

func Readlines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var results []string
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return results, nil

}
func Writeline(path string, content string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 6654)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(content + "\n"); err != nil {
		return err
	}
	return nil

}
