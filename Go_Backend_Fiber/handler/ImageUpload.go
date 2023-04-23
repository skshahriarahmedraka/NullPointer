package handler

import (
	// "bytes"
	"context"
	"fmt"
	// "io/ioutil"
	"os"

	// "mime/multipart"
	// "net/http"
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

	// f1, err := file.Open()
	// if err != nil {
	// 	return err
	// }
	// defer f1.Close()

	// fileBytes, err := ioutil.ReadAll(f1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "error ioutil.ReadAll(f1) ",
	// 	})
	// }

	// fileBuffer := bytes.NewBuffer(fileBytes)

	// contentType := http.DetectContentType(fileBytes)

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
			Region:        "us-east-1",
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

	file1, _ := file.Open()
	defer file1.Close()

	uploadInfo, err := H.MinioClient.PutObject(context.Background(), bucketName, objectName, file1, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error uploading image ",
		})
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)

	fmt.Println("Image uploaded successfully")
	imageUrl := "http://" + os.Getenv("MINIO_IP") + ":" + os.Getenv("MINIO_PORT") + "/" + bucketName + "/" + objectName
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ImageUrl": imageUrl,
	})
}
