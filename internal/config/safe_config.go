package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port       int    `json:"port" yaml:"port"`
	ServerName string `json:"servername" yaml:"servername"`
	Services   []struct {
		Name string `json:"name" yaml:"name"`
		Url  string `json:"url" yaml:"url"`
	} `json:"services" yaml:"services"`
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
		dat, err = os.ReadFile("/etc/valeera/Valeerafile")
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
