package utils

import (
	"crypto/tls"
	"errors"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

const (
	TIME_ZONE = "Asia/Jakarta"
)

var (
	DEFAULT_LANGUAGE = language.Indonesian

	ENV_PATHS = []string{"", "env"}
)

type utilsContext struct {
	config       *Config
	logger       *zap.Logger
	httpClient   *http.Client
	langBundle   *i18n.Bundle
	timeLocation *time.Location
}

var globalUtils *utilsContext

func Init() error {
	conf, err := initConfig()
	if err != nil {
		return err
	}

	var logger *zap.Logger

	// Init Configuration
	var ws zapcore.WriteSyncer = os.Stdout
	var enab zapcore.LevelEnabler = zap.DebugLevel
	if conf.Log.UdpDebugHost != "" && conf.Log.UdpInfoHost != "" && conf.Log.UdpErrorHost != "" {
		ws = &UdpDirectWriteSyncer{
			InfoHost:  conf.Log.UdpInfoHost,
			ErrorHost: conf.Log.UdpErrorHost,
			DebugHost: conf.Log.UdpDebugHost,
		}
	}

	if conf.IsProduction() {
		enab = zap.InfoLevel
	}

	// Init Logger
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, ws, enab)
	logger = zap.New(core, zap.AddCaller()).With(zap.String("service", conf.ServiceName))

	defer logger.Sync()

	// Init Language
	langBundle := initErrorLanguage()

	// init HttpClient
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// Load Time Location
	loc, err := time.LoadLocation(TIME_ZONE)
	if err != nil {
		return err
	}

	globalUtils = &utilsContext{
		config:       conf,
		logger:       logger,
		httpClient:   httpClient,
		langBundle:   langBundle,
		timeLocation: loc,
	}

	return nil
}

func GetLanguageBundle() *i18n.Bundle {
	return globalUtils.langBundle
}

func FormatDateTime(val time.Time) string {
	return val.In(GetTimeLocation()).Format(time.DateTime)
}

func GetTimeLocation() *time.Location {
	return globalUtils.timeLocation
}

func GetLogger() *zap.Logger {
	return globalUtils.logger
}

func GetConfig() *Config {
	return globalUtils.config
}

func getEnvPath() (string, error) {
	for _, path := range ENV_PATHS {
		path = filepath.Join(RootDir(), path, ".env")
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", errors.New("env is not exist")
}

func initErrorLanguage() *i18n.Bundle {
	// Initialize the i18n Bundle
	bundle := i18n.NewBundle(DEFAULT_LANGUAGE)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// Load the translation files
	bundle.MustLoadMessageFile(path.Join(RootDir(), "locales/en.yaml"))
	bundle.MustLoadMessageFile(path.Join(RootDir(), "locales/id.yaml"))

	return bundle
}
