package main

import (
	"log"
	"os"
	"pratice/udemy_questions/command_line_application/command_line_application/app"
)

func main() {

	application := app.Generate()
	
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
