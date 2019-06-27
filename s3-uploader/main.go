package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	uploadFilePath = "/Users/fukuiretu/src/github.com/fukuiretu/go-practice/s3-uploader/test-image/Gopher.png"
)

type uploader struct {
	manager   s3manager.Uploader
	src, dest string
}

func (u uploader) upload() error {
	log.Println("call upload...")

	// ファイルを読み込む
	file, err := os.Open(u.src)
	if err != nil {
		return err
	}
	defer file.Close()

	// s3にアップロード
	_, err = u.manager.Upload(&s3manager.UploadInput{
		Bucket: aws.String(u.dest),
		Key:    aws.String(filepath.Base(file.Name())),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func newUploader(s *session.Session, src, dest string) *uploader {
	return &uploader{
		manager: *s3manager.NewUploader(s),
		src:     src,
		dest:    dest,
	}
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(os.Getenv("REGION"))},
		Profile: os.Getenv("PROFILE"),
	}))
	uploader := newUploader(sess, uploadFilePath, os.Getenv("S3_BUCKET"))
	uploader.upload()
}
