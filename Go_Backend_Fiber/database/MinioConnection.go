package database

import (
	"app/logs"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)




func  MinioInit() *minio.Client {

	endpoint := os.Getenv("MINIO_IP")+":"+os.Getenv("MINIO_PORT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logs.Error("🚀 ~ file: MinioConnection.go ~ line 15 ~ funcMinioInit ~ err : ", err)
	}
    fmt.Println("😁🤩 Minio Connected successfullly : ", minioClient)
	return minioClient
}