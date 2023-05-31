package main

import (
	"fmt"

	"github.com/shahen94/floma/internal/floma"
)

func main() {

	params, err := floma.NewCommandLineParams().Parse()

	if err != nil {
		panic(err)
	}
	parser := floma.NewParser()
	config, err := parser.Parse(params.ConfigFile)

	fmt.Println(config)
	if err != nil {
		panic(err)
	}

	runner := floma.NewFlomaRunner()

	if err := runner.Run(*config); err != nil {
		panic(err)
	}
}
