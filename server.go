package main

import (
	"log"
	"template-ulamm-backend-go/api/registry"
	"template-ulamm-backend-go/api/route"
	_ "template-ulamm-backend-go/docs"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/utils"
)

//	@title			Template-Backend-ULaMM-Go
//	@version		1.0
//	@description	Template untuk inisiasi seluruh backend project pada ULaMM menggunakan bahasa pemrograman Go

//	@license.name	Apache 2.0
// 	@license.url   http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		127.0.0.1:80

func main() {
	// Init utils
	err := utils.Init()
	if err != nil {
		log.Panic(err)
	}

	// Init Datasource
	err = datasource.NewDatasource()
	if err != nil {
		log.Panic(err)
	}

	// Init Registry
	registry := registry.NewRegistry(datasource.GetDatasource())

	// Init Server
	server := route.NewServer(registry.NewController())
	if err := server.Run(":" + utils.GetConfig().PORT); err != nil {
		log.Panic(err)
	}
}
