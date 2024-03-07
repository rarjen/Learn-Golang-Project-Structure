package test

import (
	"fmt"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/utils/config"
	"testing"
)

func initConfig() {
	config.InitConfig()
}

func TestSQLServerConnection(t *testing.T) {
	initConfig()
	db, err := datasource.NewGORMDB(config.C)
	if err != nil {
		panic(err)
	}

	if pinger, ok := db.ConnPool.(interface{ Ping() error }); ok {
		if err := pinger.Ping(); err != nil {
			panic(err)
		} else {
			fmt.Println("success to ping")
		}
	} else {
		panic("failed to ping")
	}
	result := []map[string]interface{}{}
	if err := db.Raw("SELECT * FROM GolonganUser").Scan(&result).Error; err != nil {
		panic(err)
	}
	fmt.Println(result)
}
