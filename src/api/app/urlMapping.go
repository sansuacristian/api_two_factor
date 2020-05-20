package app

import (
	"github.com/gin-gonic/gin"

	"github.com/personal/api_two_factor/src/api/controller/codeGeneratorController"
	"github.com/personal/api_two_factor/src/api/controller/pingController"
	a "github.com/personal/api_two_factor/src/api/services/codeGeneratorServices"
)

//GeneratorController se usa para utilizar los metodos de las capas inferiores
type GeneratorController struct {
	code codeGeneratorController.GeneratorController
}

//BuildApplication ...
func BuildApplication() *GeneratorController {

	servicio := a.NewCodeServices()
	return &GeneratorController{code: codeGeneratorController.NewCodeController(servicio)}

}

//URLMapping inicializa las rutas para consultar en el navegador
func URLMapping(g *GeneratorController) *gin.Engine {

	router := gin.Default()
	mainRouter := router.Group("/")
	mainRouter.GET("/ping", pingController.Ping)
	mainRouter.GET("/codeGenerator", g.code.CodeGenerator)
	mainRouter.POST("/codeValidator", g.code.CodeValidator)

	return router
}
