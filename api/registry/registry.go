package registry

import (
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/pkg/datasource"
)

type Registry interface {
	NewController() controller.Controller
}

type registry struct {
	datasource *datasource.Datasource
}

func NewRegistry(datasource *datasource.Datasource) Registry {
	return &registry{
		datasource: datasource,
	}
}

func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		CommonController: r.NewCommonController(),
	}
}
