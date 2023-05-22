package objaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockery --all --case underscore --inpackage

// serviceAPI is the minimal subset of functions that we use from the AWS SDK, this allows for a greatly reduce surface
// area for mock generation.
type serviceAPI interface {
	AbortMultipartUpload(
		context.Context, *s3.AbortMultipartUploadInput, ...func(*s3.Options),
	) (*s3.AbortMultipartUploadOutput, error)

	CompleteMultipartUpload(
		context.Context, *s3.CompleteMultipartUploadInput, ...func(*s3.Options),
	) (*s3.CompleteMultipartUploadOutput, error)

	CreateMultipartUpload(
		context.Context, *s3.CreateMultipartUploadInput, ...func(*s3.Options),
	) (*s3.CreateMultipartUploadOutput, error)

	DeleteObjects(context.Context, *s3.DeleteObjectsInput, ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)
	GetObject(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	HeadObject(context.Context, *s3.HeadObjectInput, ...func(*s3.Options)) (*s3.HeadObjectOutput, error)

	ListObjectsV2Pages(
		context.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...func(*s3.Options),
	) error

	ListPartsPages(
		context.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...func(*s3.Options),
	) error

	PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	UploadPart(context.Context, *s3.UploadPartInput, ...func(*s3.Options)) (*s3.UploadPartOutput, error)

	UploadPartCopy(
		context.Context, *s3.UploadPartCopyInput, ...func(*s3.Options),
	) (*s3.UploadPartCopyOutput, error)
}
