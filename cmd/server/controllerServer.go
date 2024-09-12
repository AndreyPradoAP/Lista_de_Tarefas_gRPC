/*
	Esse documento contêm as funções que serão realizadas pelo servidor em relação ao arquivo txt
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const path = "todoList.txt"

// Function for open file
func openFile() (*os.File, error) {
	// Open/Create file
	toDoFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	return toDoFile, err
}

func serverAddTask(task string) (err error) {
	// Open/Close file ToDo
	toDoFile, err := openFile()
	if err != nil {
		panic(err)
	}
	defer toDoFile.Close()

	// Read last line in file
	scannerFile := bufio.NewScanner(toDoFile)
	var lineCount int
	var lastLine string
	for scannerFile.Scan() {
		lineCount++
		lastLine = scannerFile.Text()
	}

	if lineCount == 0 {
		// Write task in file, adding index 1
		_, err := toDoFile.WriteString("01 - " + task + "\n")
		return err
	} else {
		// Take the last index of file
		lastIndexPosition := strings.Index(lastLine, " ")
		lastIndex := lastLine[:lastIndexPosition]
		lastIndexInt, _ := strconv.Atoi(lastIndex)
		lastIndexInt++

		// Write task in file with the correct index
		if lastIndexInt < 10 {
			//Concat 0 in values less than 10
			_, err := toDoFile.WriteString("0" + strconv.Itoa(lastIndexInt) + " - " + task + "\n")
			return err
		} else {
			_, err := toDoFile.WriteString(strconv.Itoa(lastIndexInt) + " - " + task + "\n")
			return err
		}
	}
}

// Function for fulfill the task
// Precisa arrumar essa desgraça, tá apagando a lista inteira e deixando apenas a primeira linha. BURRO!
func serverDoneTask(id string) error {
	// Open file ToDo
	toDoFile, err := openFile()
	if err != nil {
		return err
	}

	// Concat 0 in values less than 10
	if id < "10" {
		id = "0" + id
	}

	// Read all lines in file and append in string
	scannerFile := bufio.NewScanner(toDoFile)
	var lines []string
	for scannerFile.Scan() {
		line := scannerFile.Text()

		// Change the line of completed task
		if strings.Contains(line, id+" - ") {
			line = strings.Replace(line, "-", "|||", 1)
		}

		lines = append(lines, line)
	}

	// Close file ToDo
	toDoFile.Close()

	// Delete the old file
	err = os.Remove(path)
	if err != nil {
		return err
	}

	// Write all tasks in a new file
	newToDoFile, err := openFile()
	if err != nil {
		return err
	}
	defer newToDoFile.Close()
	for _, line := range lines {
		_, _ = newToDoFile.WriteString(line + "\n")
	}

	return err
}

// Function for list all tasks
func serverListTasks() ([]string, error) {
	// Open/Close file
	toDoFile, err := openFile()

	// Error open file
	if err != nil {
		return []string{"error open file"}, err
	}

	defer toDoFile.Close()

	// Concat tasks in string[]
	scannerFile := bufio.NewScanner(toDoFile)
	var strTaskList []string = make([]string, 0)
	for scannerFile.Scan() {
		line := scannerFile.Text()
		strTaskList = append(strTaskList, line)
	}

	// Return for interface
	return strTaskList, nil
}

// Function for delete data in file
func serverDeleteData() error {
	toDoFile, _ := openFile()
	toDoFile.Close()

	err := os.Remove(path)

	return err
}
