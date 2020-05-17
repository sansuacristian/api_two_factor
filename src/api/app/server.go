package app

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//TestingMode Variable para saber si se encuentra en modo de testeo
var TestingMode bool

//Start metodo para que en el main levante el servidor web local
func Start() {

	constructor := BuildApplication()
	Router := URLMapping(constructor)
	StartServer(Router, TestingMode)
}

//StartServer Inicialización y configuracion del servidor
func StartServer(engine *gin.Engine, TestingMode bool) {
	if TestingMode {
		servertest := httptest.NewServer(engine)
		log.Print(servertest)
		return
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	srv.ListenAndServe()
}
