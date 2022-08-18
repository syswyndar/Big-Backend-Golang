package middleware

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/codedius/imagekit-go"
	"github.com/gin-gonic/gin"
)

func imagekitConfig() (*imagekit.Client, error) {
	opts := imagekit.Options{
		PublicKey:  os.Getenv("IMAGE_KIT_PUBLIC_KEY"),
		PrivateKey: os.Getenv("IMAGE_KIT_SECRET_KEY"),
	}

	ik, err := imagekit.NewClient(&opts)
	return ik, err
}

func uploadFile(c *gin.Context, file multipart.File, header multipart.FileHeader) (*imagekit.UploadResponse, error) {
	fileBytes, _ := ioutil.ReadAll(file)

	// // check mime type of file
	contentType := strings.Split(http.DetectContentType(fileBytes), "/")[0]
	fmt.Println(contentType, "ini content type")
	if contentType != "image" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error type of file",
		})
		c.Next()
	}
	// filename
	filename := header.Filename

	ik, err := imagekitConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	// create detail image to upload to imagekit
	ur := imagekit.UploadRequest{
		File:              fileBytes, // []byte OR *url.URL OR url.URL OR base64 string
		FileName:          filename,
		UseUniqueFileName: true,
		Tags:              []string{},
		Folder:            "/",
		IsPrivateFile:     false,
		CustomCoordinates: "",
		ResponseFields:    nil,
	}

	ctx := context.Background()

	// upload file to imagekit library
	upr, err := ik.Upload.ServerUpload(ctx, &ur)
	return upr, err
}

// middleware upload file to create product
func CreateProduct(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	upr, err := uploadFile(c, file, *header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.Set("product_img", upr.URL)
	c.Next()
}

// middleware upload file to update product
func UpdateProduct(c *gin.Context) {
	file, header, _ := c.Request.FormFile("image")
	if file == nil && header == nil {
		fmt.Println("ini error middleware")
		c.Next()
	}

	fmt.Println("harusnya tidak masuk kesini")
	upr, err := uploadFile(c, file, *header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.Set("product_img", upr.URL)
	c.Next()
}
