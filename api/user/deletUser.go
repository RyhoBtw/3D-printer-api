package user

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/api/database"
	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
)

type TemplateRequestDeletUser struct {
	Username string `form:"username"`
}

func DeletUser(c *gin.Context) {
	var request TemplateRequestNewUser
	token := c.Request.Header["Token"]
	err := c.Bind(&request)
	if err != nil {
		log.Log().Infoln(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if TestForDeletValues(request, token[0]) == true {
			DeletUserValue(request)
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "deleted successfully"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "400", "message": "user not found or Bad Request"})
		}
	}
}

func TestForDeletValues(request TemplateRequestNewUser, token string) bool {
	var user database.User
	//db := ConnectToDatabase()
	db := database.OpenDB()
	defer db.Close()

	q := fmt.Sprintf("SELECT username, loginKey FROM user WHERE username = '%s' AND loginKey = '%s'", request.Username, token)
	log.Log().Infoln(q)
	row, _ := db.Query(q)
	for row.Next() {
		_ = row.Scan(&user.Username, &user.Token)
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
	db := database.OpenDB()
	defer db.Close()

	q := fmt.Sprintf("DELETE FROM user WHERE username = '%s';", request.Username)
	//log.Log().Infoln(q)
	_, err := db.Query(q)
	if err != nil {
		log.Log().Infof("ERROT: delet %v", err)
	}
}
