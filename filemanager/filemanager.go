package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, err
	}

	file.Close()
	return lines, nil
}

func WriteJson(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil
}
