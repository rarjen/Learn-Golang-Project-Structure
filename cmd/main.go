package main

import (
	"log"
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/api/registry"
	"template-ulamm-backend-go/api/route"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/utils/config"
)

func main() {
	initConfig()
	ds, err := newDataSource(config.C)
	if err != nil {
		log.Fatal(err)
	}
	c := newController(ds)
	ginEngine := route.NewGinServer(c, config.C)
	if err := ginEngine.Run("127.0.0.1:" + config.C.PORT); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	config.InitConfig()
}

func newDataSource(conf config.Config) (*datasource.Datasource, error) {
	db, err := datasource.NewGORMDB(conf)
	if err != nil {
		return nil, err
	}
	return &datasource.Datasource{
		Db: db,
	}, nil
}

func newController(db *datasource.Datasource) controller.Controller {
	r := registry.NewRegistry(db)
	return r.NewController()
}
