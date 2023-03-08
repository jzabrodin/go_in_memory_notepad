package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const CommandExit = "exit"
const CommandCreate = "create"
const CommandList = "list"
const CommandClear = "clear"
const CommandUpdate = "update"
const CommandDelete = "delete"

var notes map[int]string
var notesSize int

func createNote(text string) {
	isNotesLimitExceeded := true

	for i := 0; i < notesSize; i++ {
		if notes[i] == "" {
			isNotesLimitExceeded = false
			break
		}
	}

	if isNotesLimitExceeded {
		fmt.Println("[Error] Notepad is full")
		return
	}

	for i := 0; i < notesSize; i++ {
		if notes[i] == "" {
			notes[i] = text
			break
		}
	}

	fmt.Println("[OK] The note was successfully created")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the maximum number of notes:")

	fmt.Scanln(&notesSize)
	notes = make(map[int]string, notesSize)

	for {
		fmt.Println("Enter a command and data:")
		scanner.Scan()
		text := scanner.Text()
		data := strings.SplitN(text, " ", 2)

		switch data[0] {
		case CommandExit:
			fmt.Println("[Info] Bye!")
			os.Exit(1)
		case CommandCreate:
			if len(data) == 1 || strings.TrimSpace(data[1]) == "" {
				fmt.Println("[Error] Missing note argument")
			} else {
				createNote(data[1])
			}
			break
		case CommandList:
			listData()
			break
		case CommandClear:
			clearMap()
			break
		case CommandUpdate:
			if len(data) == 1 || strings.TrimSpace(data[1]) == "" {
				fmt.Println("[Error] Missing position argument")
			} else {
				updateNote(data[1])
			}
			break
		case CommandDelete:
			if len(data) == 1 || strings.TrimSpace(data[1]) == "" {
				fmt.Println("[Error] Missing position argument")
			} else {
				deleteNote(data[1])
			}
			break
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func updateNote(noteData string) {

	data := strings.SplitN(noteData, " ", 2)
	i, err := strconv.ParseInt(data[0], 10, 32)

	if err != nil {
		fmt.Printf("[Error] Invalid position: %s \n", data[0])
		return
	}

	if strings.TrimSpace(noteData) == "" || strings.TrimSpace(data[0]) == "" {
		fmt.Println("[Error] Missing position argument")
		return
	}

	if len(data) < 2 || strings.TrimSpace(data[1]) == "" {
		fmt.Println("[Error] Missing note argument")
		return
	}

	if int(i) > notesSize {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", i, notesSize)
		return
	}

	if notes[int(i)-1] == "" {
		fmt.Println("[Error] There is nothing to update")
		return
	}

	notes[int(i)-1] = data[1]
	fmt.Printf("[OK] The note at position %d was successfully updated\n", i)
}

func deleteNote(noteIndexArgument string) {
	noteIndex, err := strconv.ParseInt(noteIndexArgument, 10, 64)

	if err != nil {
		fmt.Printf("[Error] Invalid position: %s \n", noteIndexArgument)
		return
	}

	if int(noteIndex) > notesSize {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", noteIndex, notesSize)
		return
	}

	realIndex := int(noteIndex) - 1
	if notes[realIndex] == "" {
		fmt.Println("[Error] There is nothing to delete")
		return
	}

	for i := realIndex; i < notesSize-1; i++ {
		notes[i] = notes[i+1]
	}

	notes[notesSize-1] = ""

	fmt.Printf("[OK] The note at position %d was successfully deleted\n", noteIndex)
}

func clearMap() {
	notes = make(map[int]string, notesSize)
	fmt.Println("[OK] All notes were successfully deleted")
}

func listData() {

	if len(notes) == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}

	for i, s := range notes {
		if s == "" {
			continue
		}
		fmt.Printf("[Info] %d: %s\n", i+1, s)
	}
}
