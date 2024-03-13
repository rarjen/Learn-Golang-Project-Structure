package main

import (
	"log"
	_ "template-ulamm-backend-go/docs"
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/api/registry"
	"template-ulamm-backend-go/api/route"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/utils/config"

)

//	@title			Template-Backend-ULaMM-Go
//	@version		1.0
//	@description	Template untuk inisiasi seluruh backend project pada ULaMM menggunakan bahasa pemrograman Go

//	@license.name	Apache 2.0
// 	@license.url   http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		127.0.0.1:80

func main() {
	initConfig()
	ds, err := newDataSource(config.C)
	if err != nil {
		log.Fatal(err)
	}
	c := newController(ds)
	ginEngine := route.NewGinServer(c, config.C)
	if err := ginEngine.Run(":" + config.C.PORT); err != nil {
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
