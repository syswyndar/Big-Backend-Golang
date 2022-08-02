package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

var (
	AWS_S3_REGION = "ap-southeast-1"
	AWS_S3_BUCKET = "ecommerce-golang"
)

func UploadImage(c *gin.Context) {
	// check req file from body data
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	filename := header.Filename
	unixtime := strconv.FormatInt(time.Now().Unix(), 10)
	unixName := unixtime + "_" + filename

	// create session
	SecretKeyID := os.Getenv("AWS_SECRET_KEY_ID")
	SecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	MyRegion := os.Getenv("AWS_REGION")
	MyBucket := os.Getenv("S3_BUCKET")
	session, sessionErr := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				SecretKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})

	// check error when create session
	if sessionErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": sessionErr,
		})
		return
	}

	//upload to the s3 bucket
	uploader := s3manager.NewUploader(session)
	up, errUpload := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(unixName),
		Body:   file,
	})

	// check error when upload file to s3
	if errUpload != nil {
		fmt.Println(errUpload, "ini error upload")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to upload file",
			"uploader": errUpload,
		})
		return
	}

	c.Set("product_img", up.Location)
	c.Next()
}
