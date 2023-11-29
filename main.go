// main.go

package main

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	_ "google.golang.org/grpc/xds/googledirectpath"
)

func main() {
	os.Setenv("STORAGE_USE_GRPC", "true")
	os.Setenv("GOOGLE_CLOUD_ENABLE_DIRECT_PATH_XDS", "true")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Read the object from bucket.
	rc, err := client.Bucket("chuangw-play").Object("test-bucket.txt").NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()
	body, err := io.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wow, the body is %s", body)
}
