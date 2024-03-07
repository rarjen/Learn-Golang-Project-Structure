package registry

import "template-ulamm-backend-go/api/controller"

func (r *registry) NewCommonController() controller.CommonController {
	return controller.NewCommonController(r.ds)
}