package app

import (
	"gin-gorm-boilerplate/internal/models"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
	lavina "github.com/Lavina-Tech-LLC/lavinagopackage"
	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {
	log := models.Logs{}

	log.Level = "Info"
	log.Msg = "Hello world"
	err := log.Log()
	if err != nil {
		llog.Error(err)
	}

	// You can call this or leave commented to return standart 200 response
	//c.JSON(200,lavina.Response("","success",true))
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) != 0 {
		c.JSON(500, lavina.Response(c.Errors, "Unexpected error", false))
	} else if c.Writer.Size() == -1 {
		c.JSON(200, lavina.Response("", "success", true))
	}
}
