package oos

import (
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

func NewUploader(client *Client, options ...func(*manager.Uploader)) *manager.Uploader {
	return manager.NewUploader(client.s3, options...)
}

func NewDownloader(client *Client, options ...func(*manager.Downloader)) *manager.Downloader {
	return manager.NewDownloader(client.s3, options...)
}
