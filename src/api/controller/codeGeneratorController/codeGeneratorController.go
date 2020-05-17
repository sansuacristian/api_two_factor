package codeGeneratorController

import (
	"net/http"

	"github.com/personal/api_two_factor/src/api/services"

	"github.com/gin-gonic/gin"
)

//GeneratorController interfaz que permite exportar los metodos que contiene es decir CodeGenerator para usarlos en otra parte del codigo
type GeneratorController interface {
	CodeGenerator(c *gin.Context)
	//CountAccess(c *gin.Context)
}

//generatorController aqui usamos la interfaz de la capa de abajo para poder usar el metodo Code
type generatorController struct {
	code services.CodeGeneratorServices
}

//NewCodeController ...
func NewCodeController(s services.CodeGeneratorServices) GeneratorController {
	return &generatorController{s}
}

//CodeGenerator le responde al usuario el codigo generado de forma automatico en la capa de servicio en el metodo code
func (g *generatorController) CodeGenerator(c *gin.Context) {

	//var userDTO dto.UserDTO

	// err := c.BindJSON(&userDTO)
	// if err != nil {
	// 	c.Status(http.StatusBadRequest)
	// }
	// log.Println("User dto", userDTO)
	response, _ := g.code.Code()
	c.JSON(http.StatusOK, gin.H{"codigo": response})
}
