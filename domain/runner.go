package domain

import "github.com/shahen94/floma/domain/parser"

type Runner interface {
	Run(parser.Config) error
}
