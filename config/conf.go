package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Conf struct {
	DatabaseHost     string `yaml:"database-host"`
	DatabaseUser     string `yaml:"database-user"`
	DatabasePassword string `yaml:"database-password"`
	JWTKey           string `yaml:"jwt-key"`
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
}

func FromYaml(dir string) (*Conf, error) {
	file, err := os.ReadFile(dir + "conf.yaml")
	if err != nil {
		panic(err)
	}
	config := Conf{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	return &config, nil
}
