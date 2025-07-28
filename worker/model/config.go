package model

type Config struct {
	Kafka struct {
		Server string `yaml:"server"`
		Group  string `yaml:"group"`
	} `yaml:"kafka"`
	Database struct {
		URL  string `yaml:"url"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}
