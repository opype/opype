package logs

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.RouterGroup){

	logs := r.Group("/logs")
	{
		logs.POST("/", processLog)
	}
}