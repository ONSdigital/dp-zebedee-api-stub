package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ONSdigital/dp-zebedee-api-stub/identity"
	"github.com/kelseyhightower/envconfig"
)

// Configuration structure which hold information for configuring the import API
type Configuration struct {
	BindAddr   string `envconfig:"BIND_ADDR"`
	Identities map[string]identity.Model
}

var cfg *Configuration

func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:   ":8082",
		Identities: make(map[string]identity.Model, 0),
	}

	b, err := ioutil.ReadFile("identity/identities.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &cfg.Identities)
	if err != nil {
		return nil, err
	}

	return cfg, envconfig.Process("", cfg)
}
