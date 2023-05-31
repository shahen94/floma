package floma

import (
	"flag"

	"github.com/shahen94/floma/domain"
)

type FlomaCommandLineParams struct{}

func (p *FlomaCommandLineParams) Parse() (domain.Params, error) {
	filePath := "floma.yml"

	flag.StringVar(&filePath, "config", filePath, "path to config file")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	return domain.Params{
		ConfigFile: filePath,
	}, nil
}

func NewCommandLineParams() domain.CommandLine {
	return &FlomaCommandLineParams{}
}
