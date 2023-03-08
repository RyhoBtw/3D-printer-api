package db

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

type TemplateRequestLogin struct {
	User     string `form:"user"`
	Passwort string `form:"passwort"`
}

type LoginResponse struct {
	Key string `form:"key"`
}

func Login(c *gin.Context) {
	var loginResponse LoginResponse
	var request TemplateRequestLogin
	err := c.Bind(&request)
	if err != nil {
		log.Log().Infoln(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if TestForLoginValues(request) == true {
		apiKey := GenerateApiKey(request, 20)

		loginResponse.Key = apiKey
		log.Log().Infof("New ApiKey generated: %s", apiKey)
		c.JSON(http.StatusOK, gin.H{"key": loginResponse.Key})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": "403", "message": "Forbidden"})
	}
	//c.Data(http.StatusOK, gin.MIMEJSON, []byte(request.Id))
}

func TestForLoginValues(request TemplateRequestLogin) bool {
	var user User
	//db := ConnectToDatabase()
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("SELECT user, passwort FROM user WHERE user = '%s' AND passwort = '%s'", request.User, request.Passwort)
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

func GenerateApiKey(request TemplateRequestLogin, n int) string {
	db := OpenDB()
	defer db.Close()
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	q := fmt.Sprintf("UPDATE user SET loginKey = '%s' WHERE user = '%s' AND passwort = '%s'", string(s), request.User, request.Passwort)
	result, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}

	rc, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("inserted %d rows\n", rc)
	return string(s)
}
