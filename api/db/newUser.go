package db

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
)

type TemplateRequestNewUser struct {
	Firstname string `form:"firstname"`
	Lastname  string `form:"lastname"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Email     string `form:"email"`
	TNR       string `form:"tnr"`
}

func NewUser(c *gin.Context) {
	var request TemplateRequestNewUser
	err := c.Bind(&request)
	if err != nil {
		log.Log().Infoln(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if CheckUsernameForNewUser(request) == true {
			InsertnewUser(request)
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "created successfully"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "username already taken"})
		}
	}
}

func InsertnewUser(request TemplateRequestNewUser) {
	//db := ConnectToDatabase()
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("INSERT INTO user (firstname, lastname, username, password, email, tnr) VALUES ('%s', '%s', '%s', '%s', '%s', '%s');", request.Firstname, request.Lastname, request.Username, request.Password, request.Email, request.TNR)
	//log.Log().Infoln(q)
	_, err := db.Query(q)
	if err != nil {
		log.Log().Infof("ERROT: Insert %v", err)
	}
}

func CheckUsernameForNewUser(request TemplateRequestNewUser) bool {
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
	if user.Username == "" {
		return true
	} else {
		return false
	}

}
