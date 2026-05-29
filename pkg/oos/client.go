package oos

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
)

type Error = smithy.APIError

type Client struct {
	s3 *s3.Client
}

func NewClient(ctx context.Context, p *profile.Profile, opts ...config.LoadOptionsFunc) (*Client, error) {
	copts := []func(*config.LoadOptions) error{
		config.WithRegion(p.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(p.AccessKey, p.SecretKey, ""),
		),
		config.WithRequestChecksumCalculation(aws.RequestChecksumCalculationWhenRequired),
		config.WithResponseChecksumValidation(aws.ResponseChecksumValidationWhenRequired),
	}
	for _, opt := range opts {
		copts = append(copts, opt)
	}
	cfg, err := config.LoadDefaultConfig(ctx, copts...)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		ep, _ := p.GetEndpoint(profile.OscServiceOOS)
		o.BaseEndpoint = &ep
		o.UsePathStyle = true
	})

	return &Client{
		s3: s3Client,
	}, nil
}

//sdk:group Bucket
func (c *Client) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	return c.s3.ListBuckets(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	return c.s3.CreateBucket(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) DeleteBucket(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOutput, error) {
	return c.s3.DeleteBucket(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	return c.s3.ListObjectsV2(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return c.s3.PutObject(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	return c.s3.DeleteObject(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	return c.s3.DeleteObjects(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error) {
	return c.s3.CopyObject(ctx, params, optFns...)
}

//sdk:group MultipartUpload
func (c *Client) AbortMultipartUpload(ctx context.Context, params *s3.AbortMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	return c.s3.AbortMultipartUpload(ctx, params, optFns...)
}

//sdk:group MultipartUpload
func (c *Client) CompleteMultipartUpload(ctx context.Context, params *s3.CompleteMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	return c.s3.CompleteMultipartUpload(ctx, params, optFns...)
}

//sdk:group MultipartUpload
func (c *Client) CreateMultipartUpload(ctx context.Context, params *s3.CreateMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	return c.s3.CreateMultipartUpload(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) DeleteBucketCors(ctx context.Context, params *s3.DeleteBucketCorsInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketCorsOutput, error) {
	return c.s3.DeleteBucketCors(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) DeleteBucketEncryption(ctx context.Context, params *s3.DeleteBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketEncryptionOutput, error) {
	return c.s3.DeleteBucketEncryption(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) DeleteBucketLifecycle(ctx context.Context, params *s3.DeleteBucketLifecycleInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketLifecycleOutput, error) {
	return c.s3.DeleteBucketLifecycle(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) DeleteBucketPolicy(ctx context.Context, params *s3.DeleteBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketPolicyOutput, error) {
	return c.s3.DeleteBucketPolicy(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) DeleteBucketWebsite(ctx context.Context, params *s3.DeleteBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketWebsiteOutput, error) {
	return c.s3.DeleteBucketWebsite(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) DeleteObjectTagging(ctx context.Context, params *s3.DeleteObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectTaggingOutput, error) {
	return c.s3.DeleteObjectTagging(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error) {
	return c.s3.GetBucketAcl(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error) {
	return c.s3.GetBucketCors(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketEncryption(ctx context.Context, params *s3.GetBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error) {
	return c.s3.GetBucketEncryption(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketLifecycleConfiguration(ctx context.Context, params *s3.GetBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	return c.s3.GetBucketLifecycleConfiguration(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketLocation(ctx context.Context, params *s3.GetBucketLocationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error) {
	return c.s3.GetBucketLocation(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketPolicy(ctx context.Context, params *s3.GetBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error) {
	return c.s3.GetBucketPolicy(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketVersioning(ctx context.Context, params *s3.GetBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error) {
	return c.s3.GetBucketVersioning(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) GetBucketWebsite(ctx context.Context, params *s3.GetBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error) {
	return c.s3.GetBucketWebsite(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return c.s3.GetObject(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) GetObjectAcl(ctx context.Context, params *s3.GetObjectAclInput, optFns ...func(*s3.Options)) (*s3.GetObjectAclOutput, error) {
	return c.s3.GetObjectAcl(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) GetObjectLockConfiguration(ctx context.Context, params *s3.GetObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error) {
	return c.s3.GetObjectLockConfiguration(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) GetObjectRetention(ctx context.Context, params *s3.GetObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error) {
	return c.s3.GetObjectRetention(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) GetObjectTagging(ctx context.Context, params *s3.GetObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error) {
	return c.s3.GetObjectTagging(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) HeadBucket(ctx context.Context, params *s3.HeadBucketInput, optFns ...func(*s3.Options)) (*s3.HeadBucketOutput, error) {
	return c.s3.HeadBucket(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	return c.s3.HeadObject(ctx, params, optFns...)
}

//sdk:group MultipartUpload
func (c *Client) ListMultipartUploads(ctx context.Context, params *s3.ListMultipartUploadsInput, optFns ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error) {
	return c.s3.ListMultipartUploads(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) ListObjectVersions(ctx context.Context, params *s3.ListObjectVersionsInput, optFns ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error) {
	return c.s3.ListObjectVersions(ctx, params, optFns...)
}

//sdk:group MultipartUpload
func (c *Client) ListParts(ctx context.Context, params *s3.ListPartsInput, optFns ...func(*s3.Options)) (*s3.ListPartsOutput, error) {
	return c.s3.ListParts(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) PutBucketAcl(ctx context.Context, params *s3.PutBucketAclInput, optFns ...func(*s3.Options)) (*s3.PutBucketAclOutput, error) {
	return c.s3.PutBucketAcl(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) PutBucketCors(ctx context.Context, params *s3.PutBucketCorsInput, optFns ...func(*s3.Options)) (*s3.PutBucketCorsOutput, error) {
	return c.s3.PutBucketCors(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) PutBucketEncryption(ctx context.Context, params *s3.PutBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.PutBucketEncryptionOutput, error) {
	return c.s3.PutBucketEncryption(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) PutBucketLifecycleConfiguration(ctx context.Context, params *s3.PutBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	return c.s3.PutBucketLifecycleConfiguration(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) PutBucketPolicy(ctx context.Context, params *s3.PutBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.PutBucketPolicyOutput, error) {
	return c.s3.PutBucketPolicy(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) PutBucketVersioning(ctx context.Context, params *s3.PutBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.PutBucketVersioningOutput, error) {
	return c.s3.PutBucketVersioning(ctx, params, optFns...)
}

//sdk:group Bucket
func (c *Client) PutBucketWebsite(ctx context.Context, params *s3.PutBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error) {
	return c.s3.PutBucketWebsite(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) PutObjectAcl(ctx context.Context, params *s3.PutObjectAclInput, optFns ...func(*s3.Options)) (*s3.PutObjectAclOutput, error) {
	return c.s3.PutObjectAcl(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) PutObjectLockConfiguration(ctx context.Context, params *s3.PutObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutObjectLockConfigurationOutput, error) {
	return c.s3.PutObjectLockConfiguration(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) PutObjectRetention(ctx context.Context, params *s3.PutObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.PutObjectRetentionOutput, error) {
	return c.s3.PutObjectRetention(ctx, params, optFns...)
}

//sdk:group Object
func (c *Client) PutObjectTagging(ctx context.Context, params *s3.PutObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error) {
	return c.s3.PutObjectTagging(ctx, params, optFns...)
}

//sdk:group MultipartUpload
func (c *Client) UploadPart(ctx context.Context, params *s3.UploadPartInput, optFns ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	return c.s3.UploadPart(ctx, params, optFns...)
}

//sdk:group MultipartUpload
func (c *Client) UploadPartCopy(ctx context.Context, params *s3.UploadPartCopyInput, optFns ...func(*s3.Options)) (*s3.UploadPartCopyOutput, error) {
	return c.s3.UploadPartCopy(ctx, params, optFns...)
}
