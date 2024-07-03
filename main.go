package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type ToDo struct {
	isDone         bool
	Text           string
	Category       string
	CompletionDate time.Time
}

func main() {
	file, err := os.Open("list.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	todos, err := ParseFile(file)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	for _, todo := range todos {
		fmt.Printf("%+v\n", todo)
	}
}

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
			}else if strings.HasPrefix(line, "-") {
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
				Text:           strings.TrimSpace(strings.TrimPrefix("F", parts[0])),
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
