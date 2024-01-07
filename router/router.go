package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	addr := "localhost:8000"
	router := gin.Default()

	initializeRoutes(router)

	router.Run(addr)
}