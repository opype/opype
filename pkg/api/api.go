package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/opype/opype/pkg/api/v1"
)

func BuildRoutes(r *gin.Engine){

	api := r.Group("/api")
	{
		v1.BuildRoutes(api)
	}
}