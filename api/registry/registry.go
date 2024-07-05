package registry

import (
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/repository"
	"template-ulamm-backend-go/pkg/usecase"
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
		CommonController: controller.NewCommonController(
			usecase.NewCommonUsecase(
				repository.NewCommonRepository(r.datasource),
			),
		),
		UserController: controller.NewUserController(
			usecase.NewUserUsecase(
				repository.NewUserRepository(r.datasource),
			),
		),
		CityController: controller.NewCityController(
			usecase.NewCityUsecase(
				repository.NewCityRepository(r.datasource),
			),
		),
	}
}
