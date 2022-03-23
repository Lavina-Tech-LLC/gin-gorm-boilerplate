package network

import (
	"net/http"

	lavina "github.com/Lavina-Tech-LLC/lavinagopackage"
	"github.com/gin-gonic/gin"
)

func NetworkRoutes(r *gin.Engine) {
	r.GET("/network/ip", ip)
}

func ip(c *gin.Context) {

	c.JSON(http.StatusOK, lavina.Response(c.Request.RemoteAddr, "", true))
}
