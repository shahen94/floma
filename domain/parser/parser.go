package parser

type Parser interface {
	Parse(path string) (*Config, error)
}
