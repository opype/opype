package logs


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func processLog(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{"message": "Log Processing Phase"})
}