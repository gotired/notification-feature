package model

type Config struct {
	Kafka struct {
		Server string `yaml:"server"`
	} `yaml:"kafka"`
	Database struct {
		URL  string `yaml:"url"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}
