package main

import (
	"fmt"
	"os"
	"time"
	"flag"

	tea "github.com/charmbracelet/bubbletea"
)

type ToDo struct {
	isDone         bool
	isHidden       bool
	Text           string
	Category       string
	Points         int
	CompletionDate time.Time
}

func main() {
	// Define flags
	help := flag.Bool("help", false, "Mostrar ayuda")
	version := flag.Bool("version", false, "Mostrar información de versión")

	// Define a command line usage description
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Madrid en 365 días.")
		fmt.Fprintln(os.Stderr, "Para ti, Kath.")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}

	// Parse flags
	flag.Parse()

	// Display help if requested
	if *help {
		flag.Usage()
		fmt.Fprint(os.Stderr, "Keys:\n")
		fmt.Fprint(os.Stderr, "\t'q': Salir de la aplicación\n")
		fmt.Fprint(os.Stderr, "\t'a': Mostrar lista completa\n")
		fmt.Fprint(os.Stderr, "\t'f': Mostrar elementos completados\n")
		fmt.Fprint(os.Stderr, "\t'r': Mostrar elementos para completar\n")
		fmt.Fprint(os.Stderr, "\n")
		os.Exit(0)
	}

	// Display version information if requested
	if *version {
		fmt.Println("Kath CLI Tool version 1.0.0")
		os.Exit(0)
	}
	listModel := initialListmodel()

	p := tea.NewProgram(listModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
