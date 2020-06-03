package server

import (
	"DormAppBackend/router"
)

// Init create a server
func Init() {
	r := router.NewRouter()
	r.Run("25.43.134.201:8080")
}
