package products

import (
	"fmt"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/api/database"
	"github.com/gin-gonic/gin"
)

type Products struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func ReturnAllProducts(c *gin.Context) {
	db := database.OpenDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var products Products
		err = rows.Scan(&products.Id, &products.Name, &products.Price)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%v\n", products)
		c.JSON(http.StatusOK, products)
		//json.NewEncoder(w).Encode(nginx)
	}
}

func JSONProduct(c *gin.Context) {

}
