package registry

import (
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/pkg/datasource"
)

type Registry interface{
	NewController() controller.Controller
}

type registry struct {
	ds *datasource.Datasource
}

func NewRegistry(ds *datasource.Datasource) Registry {
	return &registry{
		ds: ds,
	}
}

func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		CommonController: r.NewCommonController(),
		SyncDataController: r.NewSyncDataController(),
	}
}