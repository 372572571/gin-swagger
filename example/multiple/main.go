package main

import (
	swaggerFiles "github.com/372572571/files"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/372572571/gin-swagger"
	v1 "github.com/372572571/gin-swagger/example/multiple/api/v1"
	v2 "github.com/372572571/gin-swagger/example/multiple/api/v2"
	_ "github.com/372572571/gin-swagger/example/multiple/docs"
)

func main() {
	// New gin router
	router := gin.New()

	// Register api/v1 endpoints
	v1.Register(router)
	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DisplayOperationId(true),
		ginSwagger.DocExpansion("none"),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.InstanceName("v1")))

	// Register api/v2 endpoints
	v2.Register(router)
	router.GET("/swagger/v2/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName("v2")))

	// Listen and Server in
	_ = router.Run(":9582")
}
