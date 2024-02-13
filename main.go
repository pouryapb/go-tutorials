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

func main() {
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

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}

	fmt.Println("Saving todo succeeded!")

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
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
