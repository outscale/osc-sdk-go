package oos

import "github.com/aws/aws-sdk-go-v2/service/s3"

func NewPresignClient(cl *Client, opts ...func(*s3.PresignOptions)) *s3.PresignClient {
	return s3.NewPresignClient(cl.s3, opts...)
}
