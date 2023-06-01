package main

import (
	"log"

	"github.com/shahen94/floma/internal/floma"
)

func main() {

	params, err := floma.NewCommandLineParams().Parse()

	if err != nil {
		log.Fatalln(err)
	}
	parser := floma.NewParser()
	config, err := parser.Parse(params.ConfigFile)

	if err != nil {
		log.Fatalln(err)
	}

	runner := floma.NewFlomaRunner()

	if err := runner.Run(*config); err != nil {
		log.Fatalln(err)
	}
}
