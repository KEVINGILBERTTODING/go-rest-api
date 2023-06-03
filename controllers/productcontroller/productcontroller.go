package productcontroller

import (
	"net/http"

	"github.com/KEVINGILBERTTODING/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})

}
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Create(c *gin.Context) {
	namaProduct := c.PostForm("nama_product")
	deskripsi := c.PostForm("deskripsi")

	// validasi
	if namaProduct == "" || deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": "nama product dan deskripsi harus diisi"})
		return
	}

	// insert ke database
	product := models.Product{
		NamaProduct: namaProduct,
		Deskripsi:   deskripsi,
	}
	result := models.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "berhasil", "product": product})

}

func Update(c *gin.Context) {
	namaProduct := c.PostForm("nama_product")
	deskripsi := c.PostForm("deskripsi")
	id := c.PostForm("id")

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

			return
		}

	}

	product.NamaProduct = namaProduct
	product.Deskripsi = deskripsi
	models.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"status": "berhasil"})

}

func Delete(c *gin.Context) {
	id := c.PostForm("id")

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

			return
		}

	}

	result := models.DB.Delete(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "gagal", "message": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "berhasil"})

}
