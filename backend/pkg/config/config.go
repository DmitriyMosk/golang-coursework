package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Connector struct {
		ServerPort       string `yaml:"server_port"`
		JiraURL          string `yaml:"jira_url"`
		IssuesPerRequest int    `yaml:"issues_per_request"`
		NumThreads       int    `yaml:"num_threads"`
		MaxRetryWait     int    `yaml:"max_retry_wait"`
		InitialRetryWait int    `yaml:"initial_retry_wait"`
	} `yaml:"connector"`
}

func LoadConfig(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
