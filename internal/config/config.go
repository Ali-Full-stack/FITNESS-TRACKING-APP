package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      Application `yaml:"app"`
	Postgres Postgres    `yaml:"postgres"`
	Server   HTTPConfig  `yaml:"server"`
}

type Application struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"env"`
}

type HTTPConfig struct {
	Address string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout string `yaml:"timeout"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
	DBName   string `yaml:"dbname"`
}

func Load(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file")
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file")
	}

	return &config, nil
}

func (c Config) DBString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.DBName,
		c.Postgres.SSLMode,
	)
}

func (c Config) GetHostPost()string{
	return fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port)
}