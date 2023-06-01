package domain

import "github.com/shahen94/floma/domain/parser"

type Runner interface {
	UseLogger(Logger)
	Run(parser.Config) error
}
