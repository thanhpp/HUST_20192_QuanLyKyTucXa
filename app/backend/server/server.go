package server

import (
	"DormAppBackend/config"
	"DormAppBackend/router"
)

// Init create a server
func Init() {
	srvCf := config.GetServerConfig()
	r := router.NewRouter()
	r.Run(srvCf.SrvHost + ":" + srvCf.SrvPort)
}
