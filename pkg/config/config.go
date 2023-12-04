package config

import "github.com/caarlos0/env/v10"

type Config struct {
	OpenAIKey string `env:"OPEN_AI_KEY,required"`
	AIModel   string `env:"AI_MODEL,required"`
}

func New() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
