package utils

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func UploadFileToBucketEnvironment(bucketName, fileAbsolutePath, objectName string) {
	if IsRuntimeEnvironmentGCPCloudRun() || IsRuntimeEnvironmentGCPAppEngine() {
		UploadFileToBucketEnvironmentGCP(bucketName, fileAbsolutePath, objectName)
		return
	}

	panic(fmt.Errorf("no se puede subir el archivo al bucket %s en este entorno", bucketName))
}

func UploadFileToBucketEnvironmentGCP(bucketName string, fileAbsolutePath string, objectName string) {
	ctx := context.Background()

	f, err := os.Open(fileAbsolutePath)
	if err != nil {
		panic(fmt.Errorf("error abriendo el archivo %s: %w", fileAbsolutePath, err))
	}
	defer f.Close()

	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(fmt.Errorf("error creando cliente de storage: %w", err))
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	defer wc.Close()

	wc.ContentType = "application/octet-stream"

	if _, err := io.Copy(wc, f); err != nil {
		panic(fmt.Errorf("error subiendo archivo a bucket: %w", err))
	}
}
