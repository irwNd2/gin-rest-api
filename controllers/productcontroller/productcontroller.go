package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwNd2/gin-rest-api/models"
)

func Fetch(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Detail(c *gin.Context) {

}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
