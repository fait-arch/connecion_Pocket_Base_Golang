/*Server con golang*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Rutas estáticas para servir archivos estáticos
	router.Static("/static", "./public_web/static")
	router.Static("/img", "./public_web/templates/img")
	router.LoadHTMLGlob("public_web/templates/*.html")

	// Ruta para la página principal
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Habilitar el puerto 8080
	router.Run(":8080")
}
