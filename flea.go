package main

import (
	"fmt"
	"os"
)

type Project struct {
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	Contributor []string     `json:"contrib"`
	Git         string       `json:"git"`
	Depends     []Dependance `json:"depens"`
}

func New(name string) *Project {
	return &Project{
		Name:        name,
		Version:     "0.0.1",
		Contributor: []string{os.Getenv("%USERNAME%")},
		Depends:     []Dependance{Dependance{Name: "github.com/name/repo", Version: "branch Name"}}}
}

type Dependance struct {
	Name    string `json:"link"`
	Version string `json:"version"`
}

var (
	ACTION []string
)

func init() {
	if len(os.Args) > 1 {
		ACTION = append(ACTION, os.Args[1])
	} else {
		PrintHelp()
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		ACTION = append(ACTION, os.Args[2])
	}
}

func main() {
	fmt.Println(ACTION)
	switch ACTION[0] {
	case "new":
		if len(ACTION) > 1 && ACTION[1] != "" {
			NewProject(ACTION[1])
		} else {
			fmt.Printf("[!] Command 'new' need a project name \n")
		}
	case "install":
		BuildProject()
	case "save":
		SaveProject()
	default:
		return
	}
}

func PrintHelp() {
	fmt.Println("[!] No argument provided")
	fmt.Println("# Use 'new projectName' to create a project")
	fmt.Println("# Use 'install' to install project dependence")
}
