package main

import (
	"bufio"
	"strings"
	"os"
	"fmt"
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
			todo := ToDo{
				isDone:   false,
				Text:     strings.TrimSpace(strings.TrimPrefix(line, "-")),
				Category: currentCategory,
			}
			todos = append(todos, todo)
		} else if strings.HasPrefix(line, "F") {
			parts := strings.Split(line, "/")
			if len(parts) > 2 {
				return nil, fmt.Errorf("invalid format for completed task: %s", line)
			}
			dateStr := strings.TrimSpace(parts[1])
			completationDate, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid date format: %s", dateStr)
			}
			todo := ToDo{
				isDone:         true,
				Text:           strings.TrimSpace(strings.TrimPrefix(parts[0], "F")),
				Category:       currentCategory,
				CompletionDate: completationDate,
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
			str = "F "+ todo.Text + " / " + todo.CompletionDate.Format("2006-01-02") + "\n"
		} else {
			str = "- " + todo.Text + "\n"
		}
		_, err := file.WriteString(str)
		if err != nil {
			fmt.Println("Failed to write string")
			return
		}

	}
}