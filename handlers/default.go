package handlers

import (
	"lvn-tools/models"
	"net/http"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
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

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},
	)
}
