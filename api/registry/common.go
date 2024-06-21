package registry

import (
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/pkg/repository"
	"template-ulamm-backend-go/pkg/usecase"
)

func (r *registry) NewCommonController() controller.CommonController {
	return controller.NewCommonController(
		usecase.NewCommonUsecase(
			repository.NewCommonRepository(r.datasource),
		),
	)
}
