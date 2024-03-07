package config

import (
	"log"
	"os"
	"path/filepath"
	"template-ulamm-backend-go/utils"
	"template-ulamm-backend-go/utils/constantvar"

	"github.com/spf13/viper"
)

type Config struct {
	DBDSNKonve string `mapstructure:"DB_DSN_KONVE"`
	DBDSNMIS   string `mapstructure:"DB_DSN_MIS"`
	PORT       string `mapstructure:"PORT"`
	AuthAPIUrl string `mapstructure:"AUTH_API_URL"`
	STAGE      string `mapstructure:"STAGE"`
}

var C Config

func InitConfig() {
	Config := &C

	viper.AutomaticEnv()
	path := filepath.Join(utils.RootDir(), "./.env")

	// check if .env file found then read from file
	if _, err := os.Stat(path); err == nil {
		viper.SetConfigFile(path)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("failed to read config file | %v\n", err)
		}

	} else {
		viper.MustBindEnv(constantvar.DB_DSN_KONVE)
		viper.MustBindEnv(constantvar.DB_DSN_MIS)
		viper.MustBindEnv(constantvar.PORT)
		viper.MustBindEnv(constantvar.DB_DSN_MIS)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln(err)
	}

	if Config.DBDSNKonve == "" ||
		Config.PORT == "" ||
		Config.DBDSNMIS == "" ||
		Config.AuthAPIUrl == "" {
		log.Fatalln("invalid env variables")
	}
}
