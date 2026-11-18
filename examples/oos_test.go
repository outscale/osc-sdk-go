// OOS integration tests, adapted from octl's storage tests
// (octl/cmd/storage_test.go). They exercise the pkg/oos client directly
// instead of the octl CLI.
package examples_test

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/outscale/osc-sdk-go/v3/pkg/oos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const hello = "Hello world !"

// bucketName returns a unique, DNS-compliant bucket name for a test.
func bucketName(t *testing.T) string {
	t.Helper()
	sum := sha1.Sum([]byte(t.TempDir()))
	return hex.EncodeToString(sum[:])
}

func createBucket(t *testing.T, client *oos.Client, bucket string) {
	t.Helper()
	_, err := client.CreateBucket(t.Context(), &s3.CreateBucketInput{Bucket: &bucket})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, _ = client.DeleteBucket(t.Context(), &s3.DeleteBucketInput{Bucket: &bucket})
	})
}

func putObject(t *testing.T, client *oos.Client, bucket, key, body string) *s3.PutObjectOutput {
	t.Helper()
	res, err := client.PutObject(t.Context(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   strings.NewReader(body),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, _ = client.DeleteObject(t.Context(), &s3.DeleteObjectInput{Bucket: &bucket, Key: &key})
	})
	return res
}

func downloadObject(t *testing.T, client *oos.Client, bucket, key string) string {
	t.Helper()
	resp, err := client.GetObject(t.Context(), &s3.GetObjectInput{Bucket: &bucket, Key: &key})
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
	content, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(content)
}

// httpDo performs a raw HTTP request against a presigned URL and returns the
// status code and body.
func httpDo(t *testing.T, method, url, body string) (int, string) {
	t.Helper()
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	req, err := http.NewRequestWithContext(t.Context(), method, url, reader)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
	content, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return resp.StatusCode, string(content)
}

func TestOOSBucketCRUD(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)
	object := "object.txt"

	createBucket(t, client, bucket)

	res := putObject(t, client, bucket, object, hello)
	assert.NotNil(t, res.ETag)

	lres, err := client.ListObjectsV2(t.Context(), &s3.ListObjectsV2Input{Bucket: &bucket})
	require.NoError(t, err)
	require.Len(t, lres.Contents, 1)
	assert.Equal(t, object, *lres.Contents[0].Key)
}

func TestOOSObjectDownload(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	t.Run("A small object can be downloaded", func(t *testing.T) {
		small := "small.txt"
		putObject(t, client, bucket, small, hello)

		assert.Equal(t, hello, downloadObject(t, client, bucket, small))
	})

	t.Run("A large object can be downloaded", func(t *testing.T) {
		payload := strings.Repeat(hello, 100)
		large := "large.txt"
		putObject(t, client, bucket, large, payload)

		downloader := oos.NewDownloader(client)
		downloader.PartSize = 512 // force a multi-part (ranged) download
		buf := manager.NewWriteAtBuffer(nil)
		n, err := downloader.Download(t.Context(), buf, &s3.GetObjectInput{Bucket: &bucket, Key: &large})
		require.NoError(t, err)
		assert.Equal(t, int64(len(payload)), n)
		assert.Equal(t, payload, string(buf.Bytes()))
	})
}

func TestOOSBucketVersioning(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	resp, err := client.GetBucketVersioning(t.Context(), &s3.GetBucketVersioningInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.Empty(t, resp.Status)

	_, err = client.PutBucketVersioning(t.Context(), &s3.PutBucketVersioningInput{
		Bucket:                  &bucket,
		VersioningConfiguration: &types.VersioningConfiguration{Status: types.BucketVersioningStatusEnabled},
	})
	require.NoError(t, err)

	resp, err = client.GetBucketVersioning(t.Context(), &s3.GetBucketVersioningInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.Equal(t, types.BucketVersioningStatusEnabled, resp.Status)

	_, err = client.PutBucketVersioning(t.Context(), &s3.PutBucketVersioningInput{
		Bucket:                  &bucket,
		VersioningConfiguration: &types.VersioningConfiguration{Status: types.BucketVersioningStatusSuspended},
	})
	require.NoError(t, err)

	resp, err = client.GetBucketVersioning(t.Context(), &s3.GetBucketVersioningInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.Equal(t, types.BucketVersioningStatusSuspended, resp.Status)
}

func TestOOSBucketEncryption(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	_, err := client.PutBucketEncryption(t.Context(), &s3.PutBucketEncryptionInput{
		Bucket: &bucket,
		ServerSideEncryptionConfiguration: &types.ServerSideEncryptionConfiguration{
			Rules: []types.ServerSideEncryptionRule{{
				ApplyServerSideEncryptionByDefault: &types.ServerSideEncryptionByDefault{
					SSEAlgorithm: types.ServerSideEncryptionAes256,
				},
			}},
		},
	})
	require.NoError(t, err)

	resp, err := client.GetBucketEncryption(t.Context(), &s3.GetBucketEncryptionInput{Bucket: &bucket})
	require.NoError(t, err)
	require.Len(t, resp.ServerSideEncryptionConfiguration.Rules, 1)
	assert.Equal(t, types.ServerSideEncryptionAes256, resp.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm)

	_, err = client.DeleteBucketEncryption(t.Context(), &s3.DeleteBucketEncryptionInput{Bucket: &bucket})
	require.NoError(t, err)
}

func TestOOSBucketCORS(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	rules := []types.CORSRule{
		{
			AllowedHeaders: []string{"xxx", "yyy"},
			AllowedMethods: []string{"PUT", "POST", "DELETE"},
			AllowedOrigins: []string{"https://www.example.com", "https://www.foobar.example"},
			ExposeHeaders:  []string{"Content-Type"},
			MaxAgeSeconds:  aws.Int32(3000),
		},
		{
			AllowedHeaders: []string{"Authorization"},
			AllowedMethods: []string{"GET"},
			AllowedOrigins: []string{"https://yourdomain.tld", "https://www.your_domain.com"},
			MaxAgeSeconds:  aws.Int32(3000),
		},
	}

	_, err := client.GetBucketCors(t.Context(), &s3.GetBucketCorsInput{Bucket: &bucket})
	require.Error(t, err)

	_, err = client.PutBucketCors(t.Context(), &s3.PutBucketCorsInput{
		Bucket:            &bucket,
		CORSConfiguration: &types.CORSConfiguration{CORSRules: rules},
	})
	require.NoError(t, err)

	resp, err := client.GetBucketCors(t.Context(), &s3.GetBucketCorsInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.Equal(t, rules, resp.CORSRules)

	_, err = client.DeleteBucketCors(t.Context(), &s3.DeleteBucketCorsInput{Bucket: &bucket})
	require.NoError(t, err)

	_, err = client.GetBucketCors(t.Context(), &s3.GetBucketCorsInput{Bucket: &bucket})
	require.Error(t, err)
}

func TestOOSBucketPolicy(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	policy := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Deny",
            "Principal": "*",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::` + bucket + `/forbidden/*"
        }
    ]
}
`

	_, err := client.GetBucketPolicy(t.Context(), &s3.GetBucketPolicyInput{Bucket: &bucket})
	require.Error(t, err)

	_, err = client.PutBucketPolicy(t.Context(), &s3.PutBucketPolicyInput{Bucket: &bucket, Policy: &policy})
	require.NoError(t, err)

	resp, err := client.GetBucketPolicy(t.Context(), &s3.GetBucketPolicyInput{Bucket: &bucket})
	require.NoError(t, err)
	require.NotNil(t, resp.Policy)
	assert.NotEmpty(t, *resp.Policy)

	_, err = client.DeleteBucketPolicy(t.Context(), &s3.DeleteBucketPolicyInput{Bucket: &bucket})
	require.NoError(t, err)

	_, err = client.GetBucketPolicy(t.Context(), &s3.GetBucketPolicyInput{Bucket: &bucket})
	require.Error(t, err)
}

func TestOOSBucketWebsite(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	_, err := client.GetBucketWebsite(t.Context(), &s3.GetBucketWebsiteInput{Bucket: &bucket})
	require.Error(t, err)

	_, err = client.PutBucketWebsite(t.Context(), &s3.PutBucketWebsiteInput{
		Bucket: &bucket,
		WebsiteConfiguration: &types.WebsiteConfiguration{
			IndexDocument: &types.IndexDocument{Suffix: aws.String("index.html")},
		},
	})
	require.NoError(t, err)

	resp, err := client.GetBucketWebsite(t.Context(), &s3.GetBucketWebsiteInput{Bucket: &bucket})
	require.NoError(t, err)
	require.NotNil(t, resp.IndexDocument)
	assert.Equal(t, "index.html", *resp.IndexDocument.Suffix)

	_, err = client.DeleteBucketWebsite(t.Context(), &s3.DeleteBucketWebsiteInput{Bucket: &bucket})
	require.NoError(t, err)

	_, err = client.GetBucketWebsite(t.Context(), &s3.GetBucketWebsiteInput{Bucket: &bucket})
	require.Error(t, err)
}

func TestOOSBucketLifecycle(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	rules := []types.LifecycleRule{{
		ID:                          aws.String("string"),
		Status:                      types.ExpirationStatusEnabled,
		Expiration:                  &types.LifecycleExpiration{Days: aws.Int32(2)},
		NoncurrentVersionExpiration: &types.NoncurrentVersionExpiration{NoncurrentDays: aws.Int32(1)},
		Filter: &types.LifecycleRuleFilter{
			And: &types.LifecycleRuleAndOperator{
				Prefix: aws.String("string"),
				Tags:   []types.Tag{{Key: aws.String("test"), Value: aws.String("test")}},
			},
		},
	}}

	_, err := client.GetBucketLifecycleConfiguration(t.Context(), &s3.GetBucketLifecycleConfigurationInput{Bucket: &bucket})
	require.Error(t, err)

	_, err = client.PutBucketLifecycleConfiguration(t.Context(), &s3.PutBucketLifecycleConfigurationInput{
		Bucket:                 &bucket,
		LifecycleConfiguration: &types.BucketLifecycleConfiguration{Rules: rules},
	})
	require.NoError(t, err)

	resp, err := client.GetBucketLifecycleConfiguration(t.Context(), &s3.GetBucketLifecycleConfigurationInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Rules)

	_, err = client.DeleteBucketLifecycle(t.Context(), &s3.DeleteBucketLifecycleInput{Bucket: &bucket})
	require.NoError(t, err)

	_, err = client.GetBucketLifecycleConfiguration(t.Context(), &s3.GetBucketLifecycleConfigurationInput{Bucket: &bucket})
	require.Error(t, err)
}

func TestOOSBucketObjectLock(t *testing.T) {
	t.Skip("This test creates non deletable buckets, it is disabled but kept to be runned manually")

	client := newOOSClient(t)
	bucket := bucketName(t)

	_, err := client.CreateBucket(t.Context(), &s3.CreateBucketInput{
		Bucket:                     &bucket,
		ObjectLockEnabledForBucket: aws.Bool(true),
	})
	require.NoError(t, err)

	_, err = client.PutObjectLockConfiguration(t.Context(), &s3.PutObjectLockConfigurationInput{
		Bucket: &bucket,
		ObjectLockConfiguration: &types.ObjectLockConfiguration{
			ObjectLockEnabled: types.ObjectLockEnabledEnabled,
			Rule: &types.ObjectLockRule{
				DefaultRetention: &types.DefaultRetention{
					Mode: types.ObjectLockRetentionModeCompliance,
					Days: aws.Int32(1),
				},
			},
		},
	})
	require.NoError(t, err)

	resp, err := client.GetObjectLockConfiguration(t.Context(), &s3.GetObjectLockConfigurationInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.Equal(t, types.ObjectLockEnabledEnabled, resp.ObjectLockConfiguration.ObjectLockEnabled)

	object := "object.txt"
	putObject(t, client, bucket, object, hello)

	_, err = client.PutObjectRetention(t.Context(), &s3.PutObjectRetentionInput{
		Bucket: &bucket,
		Key:    &object,
		Retention: &types.ObjectLockRetention{
			Mode:            types.ObjectLockRetentionModeCompliance,
			RetainUntilDate: aws.Time(time.Now().AddDate(0, 0, 2)),
		},
	})
	require.NoError(t, err)
}

func TestOOSBucketACL(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	resp, err := client.GetBucketAcl(t.Context(), &s3.GetBucketAclInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Grants)
	before := len(resp.Grants)

	resp.Grants = append(resp.Grants, types.Grant{
		Grantee: &types.Grantee{
			Type: types.TypeGroup,
			URI:  aws.String("http://acs.amazonaws.com/groups/global/AllUsers"),
		},
		Permission: types.PermissionRead,
	})
	_, err = client.PutBucketAcl(t.Context(), &s3.PutBucketAclInput{
		Bucket:              &bucket,
		AccessControlPolicy: &types.AccessControlPolicy{Owner: resp.Owner, Grants: resp.Grants},
	})
	require.NoError(t, err)

	resp, err = client.GetBucketAcl(t.Context(), &s3.GetBucketAclInput{Bucket: &bucket})
	require.NoError(t, err)
	assert.Len(t, resp.Grants, before+1)
}

func TestOOSObjectACL(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)
	object := "object.txt"

	createBucket(t, client, bucket)
	putObject(t, client, bucket, object, hello)

	resp, err := client.GetObjectAcl(t.Context(), &s3.GetObjectAclInput{Bucket: &bucket, Key: &object})
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Grants)
	before := len(resp.Grants)

	resp.Grants = append(resp.Grants, types.Grant{
		Grantee: &types.Grantee{
			Type: types.TypeGroup,
			URI:  aws.String("http://acs.amazonaws.com/groups/global/AllUsers"),
		},
		Permission: types.PermissionRead,
	})
	_, err = client.PutObjectAcl(t.Context(), &s3.PutObjectAclInput{
		Bucket:              &bucket,
		Key:                 &object,
		AccessControlPolicy: &types.AccessControlPolicy{Owner: resp.Owner, Grants: resp.Grants},
	})
	require.NoError(t, err)

	resp, err = client.GetObjectAcl(t.Context(), &s3.GetObjectAclInput{Bucket: &bucket, Key: &object})
	require.NoError(t, err)
	assert.Len(t, resp.Grants, before+1)
}

func TestOOSObjectTagging(t *testing.T) {
	client := newOOSClient(t)
	bucket := bucketName(t)
	object := "object.txt"

	createBucket(t, client, bucket)
	putObject(t, client, bucket, object, hello)

	tags := []types.Tag{
		{Key: aws.String("key1"), Value: aws.String("value1")},
		{Key: aws.String("key2"), Value: aws.String("value2")},
	}

	resp, err := client.GetObjectTagging(t.Context(), &s3.GetObjectTaggingInput{Bucket: &bucket, Key: &object})
	require.NoError(t, err)
	assert.Empty(t, resp.TagSet)

	_, err = client.PutObjectTagging(t.Context(), &s3.PutObjectTaggingInput{
		Bucket:  &bucket,
		Key:     &object,
		Tagging: &types.Tagging{TagSet: tags},
	})
	require.NoError(t, err)

	resp, err = client.GetObjectTagging(t.Context(), &s3.GetObjectTaggingInput{Bucket: &bucket, Key: &object})
	require.NoError(t, err)
	assert.Equal(t, tags, resp.TagSet)

	_, err = client.DeleteObjectTagging(t.Context(), &s3.DeleteObjectTaggingInput{Bucket: &bucket, Key: &object})
	require.NoError(t, err)

	resp, err = client.GetObjectTagging(t.Context(), &s3.GetObjectTaggingInput{Bucket: &bucket, Key: &object})
	require.NoError(t, err)
	assert.Empty(t, resp.TagSet)
}

func TestOOSPresign(t *testing.T) {
	client := newOOSClient(t)
	presigner := oos.NewPresignClient(client)
	bucket := bucketName(t)

	createBucket(t, client, bucket)

	t.Run("Presigning GetObject works", func(t *testing.T) {
		object := "GetObject.txt"
		putObject(t, client, bucket, object, hello)

		res, err := presigner.PresignGetObject(t.Context(), &s3.GetObjectInput{Bucket: &bucket, Key: &object}, s3.WithPresignExpires(2*time.Second))
		require.NoError(t, err)
		require.True(t, strings.HasPrefix(res.URL, "http"))

		status, body := httpDo(t, http.MethodGet, res.URL, "")
		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, hello, body)

		time.Sleep(2 * time.Second)

		status, _ = httpDo(t, http.MethodGet, res.URL, "")
		assert.NotEqual(t, http.StatusOK, status)
	})

	t.Run("Presigning PutObject works", func(t *testing.T) {
		object := "PutObject.txt"
		res, err := presigner.PresignPutObject(t.Context(), &s3.PutObjectInput{Bucket: &bucket, Key: &object}, s3.WithPresignExpires(2*time.Second))
		require.NoError(t, err)
		require.True(t, strings.HasPrefix(res.URL, "http"))

		status, _ := httpDo(t, http.MethodPut, res.URL, hello)
		assert.Equal(t, http.StatusOK, status)
		t.Cleanup(func() {
			_, _ = client.DeleteObject(t.Context(), &s3.DeleteObjectInput{Bucket: &bucket, Key: &object})
		})

		assert.Equal(t, hello, downloadObject(t, client, bucket, object))
	})

	t.Run("Presigning DeleteObject works", func(t *testing.T) {
		object := "DeleteObject.txt"
		putObject(t, client, bucket, object, hello)

		res, err := presigner.PresignDeleteObject(t.Context(), &s3.DeleteObjectInput{Bucket: &bucket, Key: &object}, s3.WithPresignExpires(2*time.Second))
		require.NoError(t, err)
		require.True(t, strings.HasPrefix(res.URL, "http"))

		status, _ := httpDo(t, http.MethodDelete, res.URL, "")
		assert.Equal(t, http.StatusNoContent, status)

		_, err = client.GetObject(t.Context(), &s3.GetObjectInput{Bucket: &bucket, Key: &object})
		require.Error(t, err)
	})
}
