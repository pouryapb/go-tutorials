package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pouryapb/go-tutorials/note-app/note"
	"github.com/pouryapb/go-tutorials/note-app/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

func main() {
	printAny(4)
	printAny("interesting")

	result := add(1, 2)
	printAny(result)

	title, content := getNoteData()
	text := getUserInput("Todo Text:")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving data failed.")
		return err
	}

	fmt.Println("Saving data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func printAny(value any) {
	fmt.Println(value)
}

func add[T int | float32 | float64 | string](a, b T) T {
	return a + b
}
