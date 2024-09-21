package firebase

import (
	"context"
	_ "firebase.google.com/go"
	firebase "firebase.google.com/go"

	"fmt"
	"engine/internal/configs"

	stra "cloud.google.com/go/storage"
	"google.golang.org/api/option"
	_ "google.golang.org/api/option"
	"log"
	"net/url"
	_ "path/filepath"
)

type firebaseStore struct {
	bucket *stra.BucketHandle
}

func NewFirebaseStore(ctx context.Context, config *configs.Config) *firebaseStore {
	log.Println(config.Postgres.User)

	opt := option.WithCredentialsFile(config.FireStore.CredentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	// Get a reference to the storage client
	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("error getting storage client: %v", err)
	}

	bucket, err := client.Bucket(config.FireStore.BucketName)
	if err != nil {
		log.Fatalf("error getting bucket: %v", err)
	}
	return &firebaseStore{
		bucket: bucket,
	}
}

func (f *firebaseStore) UploadFile(ctx context.Context, fileName string, file []byte) (string, error) {

	object := f.bucket.Object(fileName)
	writer := object.NewWriter(ctx)
	if _, err := writer.Write(file); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	//Construct the public URL
	encodedPath := url.QueryEscape(fileName)
	publicURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/doppitravel-c402b.appspot.com/o/%s?alt=media", encodedPath)
	return publicURL, nil
}
