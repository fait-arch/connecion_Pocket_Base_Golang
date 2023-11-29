/*Server con golang*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//-------------------------
	//	CREACION DE DE DEPENDENCIAS POCKET BASE
	//-------------------------

	//--------------------------------------------
	//	CREACION DE UN SERVIDOR WEB SIMPLE CON GIN
	//--------------------------------------------

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Rutas estáticas para servir archivos estáticos y publicos (HTML, CSS y JavaScript)
	router.Static("/static", "./public_web/static")
	router.Static("/img", "./public_web/templates/img")
	router.LoadHTMLGlob("public_web/templates/*.html")

	// Ruta para la página principal
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Configurar la confianza en los proxies
	router.ForwardedByClientIP = true

	//
	//INCIAR POCKET BASE
	//

	//
	//INCIAR :8080
	//
	// Habilitar el puerto 8080
	fmt.Println("🅢 🅔 🅡 🅥 🅔 🅡   🅡 🅤 🅝 🏃‍♂️🏃‍♂️")
	router.Run(":8080")

}
