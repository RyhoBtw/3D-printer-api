package printer

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	cl := http.DefaultClient
	link := fmt.Sprintf("172.129.0.1:8000")
	req, _ := http.NewRequest("GET", link, nil)
	resp, err := cl.Do(req)
	if err != nil {
		log.Log().Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)
}
