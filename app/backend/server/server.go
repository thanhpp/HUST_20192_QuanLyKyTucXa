package server

import (
	"DormAppBackend/router"
)

// Init create a server
func Init() {
	r := router.NewRouter()
	r.Run("localhost:8080")
}
