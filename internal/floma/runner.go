package floma

import (
	"strconv"

	"github.com/shahen94/floma/domain"
	"github.com/shahen94/floma/domain/parser"
)

type FlomaRunner struct {
	logger     *domain.Logger
	config     *parser.Config
	processors []processorWatcher
	signal     chan parser.Command
}

type processorWatcher struct {
	processor domain.Processor
	signal    chan error
}

func (r *FlomaRunner) Run(config parser.Config) error {
	r.config = &config

	(*r.logger).WithSpinner().Info("Running floma...", "Spawning processes count: "+strconv.Itoa(len(config.Commands)))

	for _, command := range config.Commands {
		processor := NewFlomaProcessor()
		signal := make(chan error)

		r.processors = append(r.processors, processorWatcher{
			processor: processor,
			signal:    signal,
		})

		go processor.Start(command, signal)
	}

	go r.Watch()

	return nil
}

func (r *FlomaRunner) Watch() {
	for _, watcher := range r.processors {
		err := <-watcher.signal
		if err != nil {
			(*r.logger).Error(err.Error())
		} else {
			r.signal <- watcher.processor.GetCommand()
		}
	}
}

func (r *FlomaRunner) UseLogger(logger domain.Logger) {
	r.logger = &logger
}

var _ domain.Runner = &FlomaRunner{}

func NewFlomaRunner() domain.Runner {
	logger := NewLogger()
	signal := make(chan parser.Command)

	runner := &FlomaRunner{
		signal: signal,
	}

	runner.UseLogger(logger)

	return runner
}
