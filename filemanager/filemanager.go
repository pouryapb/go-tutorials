package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{inputPath, outputPath}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil
}
