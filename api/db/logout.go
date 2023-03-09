package db

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
)

type TemplateRequestLogout struct {
	Username string `form:"user"`
	LoginKey string `form:"loginKey"`
}

func Logout(c *gin.Context) {
	var request TemplateRequestLogout
	err := c.Bind(&request)
	if err != nil {
		log.Log().Infoln(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		if TestForLogoutValues(request) == true {
			DeletApiKey(request)
			c.JSON(http.StatusOK, gin.H{"info": "deleted Key"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "400", "message": "Bad Request"})
		}
	}
	//c.Data(http.StatusOK, gin.MIMEJSON, []byte(request.Id))
}

func TestForLogoutValues(request TemplateRequestLogout) bool {
	var user User
	//db := ConnectToDatabase()
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("SELECT username, loginKey FROM user WHERE username = '%s' AND loginKey = '%s'", request.Username, request.LoginKey)
	row, _ := db.Query(q)
	for row.Next() {
		_ = row.Scan(&user.Username, &user.LoginKey)
		fmt.Printf("%v\n", user)
	}
	if user.LoginKey != "" {
		return true
	} else {
		return false
	}

}

func DeletApiKey(request TemplateRequestLogout) {
	db := OpenDB()
	defer db.Close()

	q := fmt.Sprintf("UPDATE user SET loginKey = NULL WHERE username = '%s'", request.Username)
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
