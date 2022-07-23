package spec

type App struct {
	Name       string       `yaml:"name"`
	Listener   Listener     `yaml:"listener"`
	Processors []Processors `yaml:"processors"`
}

type Listener struct {
	Type string `yaml:"type"`
}

type Processors struct {
	Type string `yaml:"type"`
}
