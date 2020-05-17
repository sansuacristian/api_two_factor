package app

import (
	"github.com/gin-gonic/gin"

	"github.com/personal/api_two_factor/src/api/controller/codeGeneratorController"
	"github.com/personal/api_two_factor/src/api/controller/pingController"
	"github.com/personal/api_two_factor/src/api/services"
)

//GeneratorController se usa para utilizar los metodos de las capas inferiores
type GeneratorController struct {
	code codeGeneratorController.GeneratorController
}

//BuildApplication ...
func BuildApplication() *GeneratorController {
	servicio := services.NewCodeServices()
	return &GeneratorController{code: codeGeneratorController.NewCodeController(servicio)}

}

//URLMapping inicializa las rutas para consultar en el navegador
func URLMapping(g *GeneratorController) *gin.Engine {

	router := gin.Default()
	mainRouter := router.Group("/")
	mainRouter.GET("/ping", pingController.Ping)
	mainRouter.GET("codeGenerator", g.code.CodeGenerator)

	return router
}
