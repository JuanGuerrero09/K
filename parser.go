package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseFile(file *os.File) ([]ToDo, error) {
	var todos []ToDo
	var currentCategory string
	startParsing := false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !startParsing {
			if strings.HasPrefix(line, "###") {
				startParsing = true
			}
			continue
		}
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "@") {
			currentCategory = strings.TrimPrefix(line, "@")
		} else if strings.HasPrefix(line, "<E>") {
			currentCategory = "Bonus España"
		} else if strings.HasPrefix(line, "<EU>") {
			currentCategory = "Bonus Europa"
		} else if strings.HasPrefix(line, "-") {
			str := strings.TrimPrefix(line, "-")
			parts := strings.Split(str, "$")
			text := strings.TrimSpace(parts[0])
			points, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error not int")
				return nil, nil
			}
			todo := ToDo{
				isDone:   false,
				Text:     strings.TrimSpace(text),
				Points:   points,
				Category: currentCategory,
			}
			todos = append(todos, todo)
		} else if strings.HasPrefix(line, "F") {
			parts := strings.Split(line, "/")
			if len(parts) > 2 {
				return nil, fmt.Errorf("invalid format for completed task: %s", line)
			}
			dateStr := strings.TrimSpace(parts[1])
			subStr := strings.Split(strings.TrimPrefix(parts[0], "F"), "$")
			text := strings.TrimSpace(subStr[0])
			points, _ := strconv.Atoi(strings.TrimSpace(subStr[1]))
			completationDate, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid date format: %s", dateStr)
			}
			todo := ToDo{
				isDone:         true,
				Text:           text,
				Category:       currentCategory,
				CompletionDate: completationDate,
				Points:         points,
			}
			todos = append(todos, todo)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return todos, nil

}

func saveFile(filename string, todos []ToDo) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()
	file.WriteString("### LISTA DE PLANES EN MADRID" + "\n")
	var currentCategory string
	for _, todo := range todos {
		if currentCategory == "" {
			currentCategory = todo.Category
			file.WriteString("@" + currentCategory + "\n")
		}
		if currentCategory != todo.Category {
			currentCategory = todo.Category
			file.WriteString("@" + currentCategory + "\n")
		}
		var str string
		if todo.isDone {
			str = "F " + todo.Text + " $" + strconv.Itoa(todo.Points) + " / " + todo.CompletionDate.Format("2006-01-02") + "\n"
			} else {
			str = "- " + todo.Text + "  $" + strconv.Itoa(todo.Points) + "\n"
		}
		_, err := file.WriteString(str)
		if err != nil {
			fmt.Println("Failed to write string")
			return
		}

	}
}
