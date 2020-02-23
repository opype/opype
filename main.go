package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opype/opype/pkg/api"
)

func main() {

	router := gin.Default()
	api.BuildRoutes(router)
	_ = router.Run(fmt.Sprintf(":%d", 8080))
}
