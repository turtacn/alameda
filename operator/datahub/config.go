package datahub

import (
	"errors"
	"net/url"

	datahubutils "github.com/turtacn/alameda/operator/pkg/utils/datahub"
)

type retryInterval struct {
	Default int64 `mapstructure:"default"`
}

type Config struct {
	Address       string        `mapstructure:"address"`
	RetryInterval retryInterval `mapstructure:"retryInterval"`
}

func NewConfig() *Config {

	c := Config{}
	c.init()
	return &c
}

func (c *Config) init() {
	c.Address = datahubutils.GetDatahubAddress()
	c.RetryInterval = retryInterval{
		Default: 3,
	}
}

func (c *Config) Validate() error {

	var err error

	_, err = url.Parse(c.Address)
	if err != nil {
		return errors.New("datahub config validate failed: " + err.Error())
	}

	return nil
}
