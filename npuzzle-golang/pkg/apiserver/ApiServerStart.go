package apiserver

import (
	handler "N-puzzle-GO/api"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiServerStart() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AddAllowMethods("OPTIONS", "PUT", "DELETE", "GET")
	router.Use(cors.New(corsConfig), gin.Recovery())
	handler.RequestHandler(router)
	if err := router.Run(); err != nil {
		fmt.Errorf("%s", err)
	}
}
