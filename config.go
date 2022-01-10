package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const configDir = ".hhlog"

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

func parseConfig(data []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func readConfig() (config *Config) {
	userHomeDir, err := os.UserHomeDir()
	var data []byte
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find user home directory: %v\n", err)
		fmt.Fprintf(os.Stderr, "Trying config file in the working directory\n")
		goto workingDir
	}
	data, err = ioutil.ReadFile(userHomeDir + "/" + configDir + "/hhlog.conf")
	if err == nil {
		goto parse
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config from home dir: %v\n", err)
		fmt.Fprintf(os.Stderr, "Trying config file in the working directory\n")
	}
workingDir:
	data, err = ioutil.ReadFile(".hhlog.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config file:\n")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "Proceeding without config file.\n")
		return nil
	}
parse:
	config, err = parseConfig(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse config file:\n")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "Proceeding without config file.\n")
		config = nil
	}
	return
}
