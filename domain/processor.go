package domain

import "github.com/shahen94/floma/domain/parser"

type Processor interface {
	Start(parser.Command, chan<- error)
	GetCommand() parser.Command
	Stop()
}
