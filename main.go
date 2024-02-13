package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pouryapb/go-tutorials/note-app/note"
	"github.com/pouryapb/go-tutorials/note-app/todo"
)

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
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving todo failed.")
		return
	}

	fmt.Println("Saving todo succeeded!")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving note failed.")
		return
	}

	fmt.Println("Saving note succeeded!")
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
