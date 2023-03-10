package user

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/api/database"
	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

type TemplateRequestLogin struct {
	Username string `form:"username"`
	Password string `form:"password"`
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
	var user database.User
	//db := ConnectToDatabase()
	db := database.OpenDB()
	defer db.Close()

	q := fmt.Sprintf("SELECT username, password FROM user WHERE username = '%s' AND password = '%s'", request.Username, request.Password)
	//log.Log().Infoln(q)
	row, _ := db.Query(q)
	for row.Next() {
		_ = row.Scan(&user.Username, &user.Password)
		fmt.Printf("%v\n", user)
	}
	if user.Username != "" {
		return true
	} else {
		return false
	}

}

func GenerateApiKey(request TemplateRequestLogin, n int) string {
	db := database.OpenDB()
	defer db.Close()
	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]byte, n)
	for i := range s {
		// randomly select 1 character from given charset
		s[i] = charset[rand.Intn(len(charset))]
	}

	q := fmt.Sprintf("UPDATE user SET loginKey = '%s' WHERE username = '%s' AND password = '%s'", string(s), request.Username, request.Password)
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
