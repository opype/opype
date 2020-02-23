package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opype/opype/pkg/api/v1/logs"
	"net/http"
)

func ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func BuildRoutes(r *gin.RouterGroup){

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
		logs.BuildRoutes(v1)
	}
}