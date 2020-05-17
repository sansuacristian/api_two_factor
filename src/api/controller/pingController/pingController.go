package pingController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//PingController necesario para exportar los metodos que incorpora
type PingController interface {
	Ping(c *gin.Context)
}

//Ping respuesta que se le va a dar al usuario en el navegador web
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "pong"})
}
