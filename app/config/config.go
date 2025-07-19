package config

import (
	"log"
	"os"

	"github.com/gotired/notification-feature/app/model"
	"gopkg.in/yaml.v2"
)

func Load(path string) *model.Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}
	defer file.Close()

	cfg := &model.Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)
	if err != nil {
		log.Fatalf("error decode YAML: %v", err)
	}
	return cfg
}
