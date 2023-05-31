package floma

import (
	"os"

	"github.com/shahen94/floma/domain/parser"
	"gopkg.in/yaml.v3"
)

type FlomaParser struct{}

func (p *FlomaParser) Parse(path string) (*parser.Config, error) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	config := &parser.Config{}

	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func NewParser() parser.Parser {
	return &FlomaParser{}
}
