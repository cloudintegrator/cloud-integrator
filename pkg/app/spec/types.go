package spec

type AppSpec struct {
	Name       string           `yaml:"name"`
	Listener   ListenerSpec     `yaml:"listener"`
	Processors []ProcessorsSpec `yaml:"processors"`
}

type ListenerSpec struct {
	Type   string `yaml:"type"`
	Config Config `yaml:"config"`
}

type ProcessorsSpec struct {
	Type   string `yaml:"type"`
	Config Config `yaml:"config"`
}

type Config struct {
}
