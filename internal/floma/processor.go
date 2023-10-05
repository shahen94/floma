package floma

import (
	"os"
	"os/exec"

	"github.com/shahen94/floma/domain"
	"github.com/shahen94/floma/domain/parser"
)

// FlomaProcessor is a struct that implements domain.Processor interface.
type FlomaProcessor struct {
	Command parser.Command
	process *os.Process
	signal  chan<- error
}

// execute executes the command.
func (p *FlomaProcessor) execute() {
	path, err := os.Getwd()

	if err != nil {
		p.signal <- err
		return
	}

	if p.Command.Dir != nil {
		path = *p.Command.Dir
	}

	cmd := exec.Cmd{
		Path: path,
		Args: p.Command.Args,
		Env:  p.Command.Environment,
	}

	p.process = cmd.Process

	err = cmd.Start()

	p.signal <- err
}

// GetCommand returns the command.
func (p *FlomaProcessor) GetCommand() parser.Command {
	return p.Command
}

// Start starts the command.
func (p *FlomaProcessor) Start(command parser.Command, signal chan<- error) {
	p.Command = command
	p.signal = signal

	p.execute()
}

// Stop stops the command.
func (p *FlomaProcessor) Stop() {
	if p.process != nil {
		p.process.Kill()
	}
	p.signal <- nil
}

// NewFlomaProcessor returns a new FlomaProcessor.
func NewFlomaProcessor() domain.Processor {
	return &FlomaProcessor{}
}
