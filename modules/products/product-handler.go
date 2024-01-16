package products

import (
	"fmt"
	"golang-basic/responses"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	CreateProduct(c *gin.Context)
	GetAllProducts(c *gin.Context)
	GetProductById(c *gin.Context)
	UpdateProductById(c *gin.Context)
	DeleteProductById(c *gin.Context)
}

type productHandler struct {
	productService IProductService
}

func NewProductHandler(productService IProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var product Products

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err.Error()))
		return
	}
	image, _ := DecodeAndSaveImage(product.Image)
	fmt.Println("image")
	fmt.Println(image)

	createdProduct, err := h.productService.CreateProduct(c.Request.Context(), product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating product"})
		return
	}

	c.JSON(http.StatusCreated, responses.NewSuccessResponse("Success Create Products", createdProduct))

}

func DecodeAndSaveImage(base64Image string) (string, error) {
	return base64Image, nil
}

// Add a function to save the file to disk
func SaveFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (h *productHandler) GetAllProducts(c *gin.Context) {
	var products []Products
	products, err := h.productService.GetAllProduct(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating product"})
		return
	}
	c.JSON(http.StatusOK, responses.NewSuccessResponse("Success Get All Products", products))
}

func (h *productHandler) GetProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productService.GetProductById(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get product"})
		return
	}
	c.JSON(http.StatusOK, responses.NewSuccessResponse("Success Get Products", product))
}

func (h *productHandler) UpdateProductById(c *gin.Context) {
	var product Products
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	updatedProduct, err := h.productService.UpdateProductById(c.Request.Context(), uint(id), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get product"})
		return
	}
	c.JSON(http.StatusOK, responses.NewSuccessResponse("Success Get Update Products", updatedProduct))
}

func (h *productHandler) DeleteProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productService.DeleteProductById(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get product"})
		return
	}
	c.JSON(http.StatusOK, responses.NewSuccessResponse("Success Delete Products", product))
}
