package architecture

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port     int `yaml:"port" default:"8080"`
	Database struct {
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		Database     string `yaml:"database"`
		Host         string `yaml:"host"`
		MigrationSrc string `yaml:"migration_src"`
	} `yaml:"db"`
}

func LoadConfig() (Config, error) {
	configFileName := flag.String("config", "../../config.yml", "Specify config file")
	flag.Parse()
	file, err := os.ReadFile(*configFileName)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
