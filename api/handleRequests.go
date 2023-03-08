package api

import (
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/api/db"
	"github.com/RyhoBtw/3D-printer-api/api/printer"
	"github.com/gin-gonic/gin"
)

// handle api requests
func HandleRequests() {
	r := gin.Default()
	g := r.Group("/api/v1")
	{
		g.GET("/print", HomePage)
		g.GET("/print/status", printer.GetStatus)
		g.GET("/print/login", db.Login)
		//g.POST("/print/Gcode", printer.PostGcode)
	}
	r.Run()
}

func HomePage(c *gin.Context) {
	//jsonData := []byte(`404 Not Found`)
	c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	//c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
}
