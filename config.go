package main

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Station Station `yaml:"station"`
	Pota    Pota    `yaml:"pota"`
	Wwff    Wwff    `yaml:"wwff"`
}

type Station struct {
	Call string `yaml:"call"`
	Mail Mail   `yaml:"mail"`
}

type Pota struct {
	ContactName  string `yaml:"name"`
	ContactEmail string `yaml:"email"`
}

type Wwff struct {
	ContactName  string `yaml:"name"`
	ContactEmail string `yaml:"email"`
}

type Mail struct {
	SmtpHost string `yaml:"smtp-host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Email    string `yaml:"email"`
}

func readConfig(data []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
