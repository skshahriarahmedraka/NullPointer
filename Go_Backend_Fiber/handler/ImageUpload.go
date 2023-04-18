package handler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"

	// "mime/multipart"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/minio/minio-go/pkg/credentials"
	"github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/google/uuid"
)

func (H *DatabaseCollections) UploadImage(c *fiber.Ctx) error {
	c.Accepts("application/json")

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid file",
		})
	}

	f1, err := file.Open()
	if err != nil {
		return err
	}
	defer f1.Close()


	fileBytes, err := ioutil.ReadAll(f1)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error ioutil.ReadAll(f1) ",
		})
	}


	fileBuffer := bytes.NewBuffer(fileBytes)

	
	contentType := http.DetectContentType(fileBytes)

	// endpoint := "0.0.0.0:9000"
	// accessKey := "SZeY42LO5pm6I0l7"
	// secretKey := "aUqCSe9lZs1Auh0sCETp9sVD9INGH1CX"
	// useSSL := false
	// minioClient, err := minio.New(endpoint, &minio.Options{
	// 	Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
	// 	Secure: useSSL,
	// })
    // fmt.Println("🚀 ~ file: UploadImage.go ~ line 114 ~ func ~ minioClient : ", minioClient)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Minio connection error",
	// 	})
	// }


	bucketName := "nullpointerbucket"
	objectName := uuid.New().String()


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	exists, err := H.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
    fmt.Println("🚀 ~  minioClient.BucketExists(ctx, bucketName) : ", err)
		fmt.Println(err)

	}
	if !exists {
		opts := minio.MakeBucketOptions{
			Region: "us-east-1",
			ObjectLocking: true,
		}
		err = H.MinioClient.MakeBucket(ctx, bucketName, opts)
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println("Bucket created successfully")
	} else {
		fmt.Println("Bucket already exists")
	}


	uploadInfo, err := H.MinioClient.PutObject(context.Background(), bucketName, objectName, fileBuffer, int64(len(fileBytes)), minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error uploading image ",
		})
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)

	fmt.Println("Image uploaded successfully")
	imageUrl := "http://"+ os.Getenv("MINIO_IP") +":"+ os.Getenv("MINIO_PORT") +"/" +bucketName+"/" +objectName
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ImageUrl": imageUrl,
	})
}

