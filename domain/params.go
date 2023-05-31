package domain

type Params struct {
	ConfigFile string
}

type CommandLine interface {
	Parse() (Params, error)
}
