package user

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/api/database"
	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
)

type TemplateRequestLogout struct {
	Username string `form:"username"`
}

func Logout(c *gin.Context) {
	var request TemplateRequestLogout
	token := c.Request.Header["Token"]
	err := c.Bind(&request)
	if err != nil {
		log.Log().Infoln(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		if TestForLogoutValues(request, token[0]) == true {
			DeletApiKey(request)
			c.JSON(http.StatusOK, gin.H{"info": "deleted Key"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "400", "message": "Bad Request"})
		}
	}
	//c.Data(http.StatusOK, gin.MIMEJSON, []byte(request.Id))
}

func TestForLogoutValues(request TemplateRequestLogout, token string) bool {
	var user database.User
	db := database.OpenDB()
	defer db.Close()

	q := fmt.Sprintf("SELECT username, loginKey FROM user WHERE username = '%s' AND loginKey = '%s'", request.Username, token)
	row, _ := db.Query(q)
	for row.Next() {
		_ = row.Scan(&user.Username, &user.Token)
		fmt.Printf("%v\n", user)
	}
	if user.Token != "" {
		return true
	} else {
		return false
	}

}

func DeletApiKey(request TemplateRequestLogout) {
	db := database.OpenDB()
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
