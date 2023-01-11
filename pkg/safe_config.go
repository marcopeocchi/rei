package pkg

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type SafeConfig struct {
	mu  sync.Mutex
	cfg Config
}

// TODO
func (c *SafeConfig) Save(value Config) {
	c.mu.Lock()
	c.cfg = value
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
	yaml.Unmarshal(dat, &c.cfg)
}

func (c *SafeConfig) Json() []byte {
	bytes, err := json.Marshal(c.cfg)
	if err != nil {
		log.Panicln(err)
	}
	return bytes
}
