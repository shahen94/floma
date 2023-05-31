package parser

type Config struct {
	Commands []Command `yaml:"commands"`
}

type Command struct {
	Name        *string  `yaml:"name"`
	Exec        string   `yaml:"exec"`
	Args        []string `yaml:"args"`
	Environment []string `yaml:"environment"`
	Dir         *string  `yaml:"dir"`
}
