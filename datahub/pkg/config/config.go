package config

import (
	"errors"
	Keycodes "github.com/turtacn/alameda/datahub/pkg/account-mgt/keycodes"
	Apis "github.com/turtacn/alameda/datahub/pkg/apis"
	Notifier "github.com/turtacn/alameda/datahub/pkg/notifier"
	InternalInflux "github.com/turtacn/alameda/internal/pkg/database/influxdb"
	InternalLdap "github.com/turtacn/alameda/internal/pkg/database/ldap"
	InternalPromth "github.com/turtacn/alameda/internal/pkg/database/prometheus"
	InternalRabbitMQ "github.com/turtacn/alameda/internal/pkg/message-queue/rabbitmq"
	InternalWeaveScope "github.com/turtacn/alameda/internal/pkg/weavescope"
	"github.com/turtacn/alameda/pkg/utils/log"
)

const (
	defaultBindAddress = ":50050"
)

type Config struct {
	BindAddress string `mapstructure:"bindAddress"`
	ClusterUID  string
	Prometheus  *InternalPromth.Config     `mapstructure:"prometheus"`
	InfluxDB    *InternalInflux.Config     `mapstructure:"influxdb"`
	Apis        *Apis.Config               `mapstructure:"apis"`
	Ldap        *InternalLdap.Config       `mapstructure:"ldap"`
	Keycode     *Keycodes.Config           `mapstructure:"keycode"`
	Notifier    *Notifier.Config           `mapstructure:"notifier"`
	WeaveScope  *InternalWeaveScope.Config `mapstructure:"weavescope"`
	RabbitMQ    *InternalRabbitMQ.Config   `mapstructure:"rabbitmq"`
	Log         *log.Config                `mapstructure:"log"`
}

func NewDefaultConfig() Config {
	var (
		defaultLogConfig        = log.NewDefaultConfig()
		defaultPrometheusConfig = InternalPromth.NewDefaultConfig()
		defaultInfluxDBConfig   = InternalInflux.NewDefaultConfig()
		defaultApisConfig       = Apis.NewDefaultConfig()
		defaultLdapConfig       = InternalLdap.NewDefaultConfig()
		defaultKeycodeConfig    = Keycodes.NewDefaultConfig()
		defaultNotifierConfig   = Notifier.NewDefaultConfig()
		defaultWeaveScopeConfig = InternalWeaveScope.NewDefaultConfig()
		defaultRabbitMQConfig   = InternalRabbitMQ.NewDefaultConfig()
		config                  = Config{
			BindAddress: defaultBindAddress,
			Prometheus:  defaultPrometheusConfig,
			InfluxDB:    defaultInfluxDBConfig,
			Apis:        defaultApisConfig,
			Ldap:        defaultLdapConfig,
			Keycode:     defaultKeycodeConfig,
			Notifier:    defaultNotifierConfig,
			WeaveScope:  defaultWeaveScopeConfig,
			RabbitMQ:    defaultRabbitMQConfig,
			Log:         &defaultLogConfig,
		}
	)

	defaultKeycodeConfig.InfluxDB = defaultInfluxDBConfig
	defaultKeycodeConfig.Ldap = nil // TODO: defaultLdapConfig

	return config
}

func (c *Config) Validate() error {
	var err error

	err = c.Prometheus.Validate()
	if err != nil {
		return errors.New("failed to validate gRPC config: " + err.Error())
	}

	return nil
}
