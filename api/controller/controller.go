package controller

type Controller struct {
	CommonController   interface{ CommonController }
	SyncDataController interface{ SyncDataController }
}
