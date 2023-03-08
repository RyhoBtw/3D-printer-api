package db

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
)

type TemplateRequestLogout struct {
	User     string `form:"user"`
	Passwort string `form:"passwort"`
	LoginKey string `form:"loginKey`
}

func Logout(c *gin.Context) {
	var request TemplateRequestLogout
	err := c.Bind(&request)
	if err != nil {
		log.Log().Infoln(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if TestForLogin(request) == true {
		DeletApiKey(request)
		c.JSON(http.StatusOK, gin.H{"info": "deleted Key"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": "400", "message": "Bad Request"})
	}
	//c.Data(http.StatusOK, gin.MIMEJSON, []byte(request.Id))
}

func TestForLogin(request TemplateRequestLogout) bool {
	var user User
	//db := ConnectToDatabase()
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("SELECT user, passwort, loginKey FROM user WHERE user = '%s' AND passwort = '%s' AND loginKey = '%s'", request.User, request.Passwort, request.LoginKey)
	//log.Log().Infoln(q)
	row, _ := db.Query(q)
	for row.Next() {
		_ = row.Scan(&user.User, &user.Passwort)
		fmt.Printf("%v\n", user)
	}
	if user.User != "" {
		return true
	} else {
		return false
	}

}

func DeletApiKey(request TemplateRequestLogout) {
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("UPDATE user SET loginKey = NULL WHERE user = '%s' AND passwort = '%s'", request.User, request.Passwort)
	result, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}

	rc, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("inserted %d rows\n", rc)
}
