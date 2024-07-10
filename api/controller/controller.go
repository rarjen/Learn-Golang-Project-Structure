package controller

import "time"

const (
	TIMEOUT = time.Second * 15
)

type Controller struct {
	CommonController  interface{ CommonController }
	UserController    interface{ UserController }
	CityController    interface{ CityController }
	ProgramController interface{ ProgramController }
	ProductController interface{ ProductController }
}
