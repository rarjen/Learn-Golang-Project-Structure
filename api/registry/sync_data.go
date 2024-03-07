package registry

import (
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/pkg/repository"
	"template-ulamm-backend-go/pkg/usecase"
)

func (r *registry) NewSyncDataController() controller.SyncDataController {
	externalRepo := repository.NewExternalSourceRepository(r.ds)
	syncDataRepo := repository.NewSyncDataRepository(r.ds, externalRepo)
	globalRepo := repository.NewGlobalRepository()
	syncDataUC := usecase.NewSyncDataUseCase(syncDataRepo, globalRepo)
	return controller.NewSyncDataController(syncDataUC)
}