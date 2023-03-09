package db

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
)

type TemplateRequestDeletUser struct {
	Username string `form:"username"`
	LoginKey string `form:"loginKey"`
}

func DeletUser(c *gin.Context) {
	var request TemplateRequestNewUser
	err := c.Bind(&request)
	if err != nil {
		log.Log().Infoln(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if CheckUsernameForDelet(request) == true {
			DeletUserValue(request)
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "deleted successfully"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "user not found"})
		}
	}
}

func CheckUsernameForDelet(request TemplateRequestNewUser) bool {
	var user User
	//db := ConnectToDatabase()
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("SELECT username FROM user WHERE username = '%s'", request.Username)
	//log.Log().Infoln(q)
	row, _ := db.Query(q)
	for row.Next() {
		_ = row.Scan(&user.Username)
		fmt.Printf("%v\n", user)
	}
	if user.Username != "" {
		return true
	} else {
		return false
	}

}

func DeletUserValue(request TemplateRequestNewUser) {
	//db := ConnectToDatabase()
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("DELETE FROM user WHERE username = '%s';", request.Username)
	//log.Log().Infoln(q)
	_, err := db.Query(q)
	if err != nil {
		log.Log().Infof("ERROT: delet %v", err)
	}
}
