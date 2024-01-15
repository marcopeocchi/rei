package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Url         string `json:"url" yaml:"url"`
	Name        string `json:"name" yaml:"name"`
	Icon        string `json:"icon" yaml:"icon"`
	Description string `json:"description" yaml:"description"`
}

type Config struct {
	Port           int       `json:"port" yaml:"port"`
	ServerName     string    `json:"servername" yaml:"servername"`
	Authentication bool      `json:"authentication" yaml:"authentication"`
	Username       string    `json:"username" yaml:"username"`
	Password       string    `json:"password" yaml:"password"`
	Services       []Service `json:"services" yaml:"services"`
	Arr            []Service `json:"arr" yaml:"arr"`
	Media          []Service `json:"media" yaml:"media"`
	Downloaders    []Service `json:"downloaders" yaml:"downloaders"`
	System         []Service `json:"system" yaml:"system"`
	Devices        []Service `json:"devices" yaml:"devices"`
}

type SafeConfig struct {
	mu  sync.Mutex
	Cfg Config
}

func New(path string) *SafeConfig {
	c := &SafeConfig{}
	c.Load(path)
	return c
}

// TODO
func (c *SafeConfig) Save(value Config) {
	c.mu.Lock()
	c.Cfg = value
	c.mu.Unlock()
}

func (c *SafeConfig) Load(configPath string) {
	dat, err := os.ReadFile(configPath)
	if err != nil {
		dat, err = os.ReadFile("/etc/rei/config.yml")
		if err != nil {
			log.Fatalln(err)
		}
	}
	yaml.Unmarshal(dat, &c.Cfg)
}

// Returns JSON representation on Config as it is initially saved as YAML
func (c *SafeConfig) Json() []byte {
	bytes, err := json.Marshal(c.Cfg)
	if err != nil {
		log.Panicln(err)
	}
	return bytes
}

// Returns JSON representation on Config as it is initially saved as YAML
func (c *SafeConfig) JsonEncoder(w io.Writer) error {
	return json.NewEncoder(w).Encode(c.Cfg)
}
