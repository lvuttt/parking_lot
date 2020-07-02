package main

import (
	"os"
	"parking_lot/internal/app/modeler"
)

func main() {
	var mode modeler.Modeler
	args := os.Args[1:]
	if len(args) == 0 {
		mode = &modeler.Interaction{}
	} else {
		mode = &modeler.File{
			FilePath: args[0],
		}
	}
	mode.Start()
}
