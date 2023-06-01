package floma

import (
	"strconv"

	"github.com/shahen94/floma/domain"
	"github.com/shahen94/floma/domain/parser"
)

type FlomaRunner struct {
	logger *domain.Logger
	config *parser.Config
}

func (r *FlomaRunner) Run(config parser.Config) error {
	r.config = &config

	(*r.logger).WithSpinner().Info("Running floma...", "Spawning processes count: "+strconv.Itoa(len(config.Commands)))

	return nil
}

func (r *FlomaRunner) UseLogger(logger domain.Logger) {
	r.logger = &logger
}

var _ domain.Runner = &FlomaRunner{}

func NewFlomaRunner() domain.Runner {
	logger := NewLogger()
	runner := &FlomaRunner{}

	runner.UseLogger(logger)

	return runner
}
