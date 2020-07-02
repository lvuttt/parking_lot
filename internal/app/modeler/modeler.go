package modeler

import (
	"fmt"
	"log"
	"strings"
	"os"
	"bufio"
	"parking_lot/internal/app/translate"
)

type (
	File struct {
		FilePath string
		Lines []string
	}
	Interaction struct {
	}
)

func (f *File) Start() {
	lines, err := f.readLines()
	if err != nil {
		log.Fatalf("cannot read files :%s\n", err)
	}
	commands := [][]string{}
	for _, line := range lines {
		commands = append(commands, strings.Split(line, " "))
	}
	tran := &translate.Translation{Commands: commands}
	tran.Run()
	fmt.Print(strings.Join(tran.Park.Stdout, ""))
}

func (f *File) readLines() ([]string, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (i *Interaction) Start() {
	tran := &translate.Translation{}
	for {
		fmt.Print("$ ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			if line == "exit" {
				break
			}
			command := strings.Split(strings.TrimSuffix(line, "\\n"), " ")
			_, err := tran.Operate(command)
			fmt.Print(tran.Park.Stdout[len(tran.Park.Stdout)-1])
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		}
	}
}