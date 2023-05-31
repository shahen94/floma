package floma

import (
	"github.com/shahen94/floma/domain"
	"github.com/shahen94/floma/domain/parser"
)

type FlomaRunner struct{}

func (r *FlomaRunner) Run(config parser.Config) error {
	return nil
}

func NewFlomaRunner() domain.Runner {
	return &FlomaRunner{}
}
