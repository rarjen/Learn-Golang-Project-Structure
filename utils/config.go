package utils

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"template-ulamm-backend-go/pkg/constantvar"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type (
	Config struct {
		PORT        string `mapstructure:"PORT" validate:"required"`
		STAGE       string `mapstructure:"STAGE" validate:"required"`
		ServiceName string `mapstructure:"SERVICE_NAME" validate:"required"`

		SqlServer SqlServer `mapstructure:",squash"`
		Log       Log       `mapstructure:",squash"`
		Redis     Redis     `mapstructure:",squash"`
		Mongo     Mongo     `mapstructure:",squash"`
	}

	SqlServer struct {
		Host  string `mapstructure:"SQL_SERVER_URI" validate:"required"`
		Hosts []string
	}

	Log struct {
		UdpInfoHost  string `mapstructure:"LOG_UDP_INFO_HOST"`
		UdpErrorHost string `mapstructure:"LOG_UDP_ERROR_HOST"`
		UdpDebugHost string `mapstructure:"LOG_UDP_DEBUG_HOST"`
	}

	Mongo struct {
		Host   string `mapstructure:"MONGODB_URI"`
		DbName string `mapstructure:"MONGODB_DB_NAME"`
	}

	Redis struct {
		Host     string `mapstructure:"REDIS_URI"`
		Password string `mapstructure:"REDIS_PASSWORD"`
	}
)

func initConfig() (*Config, error) {
	viper.AutomaticEnv()

	path, err := getEnvPath()
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file | %v", err)
	}

	var conf *Config
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(conf); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, fmt.Errorf("invalid config:%v", err)
		}

		return nil, err
	}

	conf.SqlServer.Hosts = strings.Split(conf.SqlServer.Host, "|")

	return conf, nil
}

// Fetch root directory path
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func (conf *Config) IsDevelopment() bool {
	return conf.STAGE == constantvar.CONFIG_STAGE_DEV
}

func (conf *Config) IsProduction() bool {
	return conf.STAGE == constantvar.CONFIG_STAGE_PROD
}
