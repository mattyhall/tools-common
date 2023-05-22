package objaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type OptionFn func(*s3.Options)

//go:generate mockery --all --case underscore --inpackage

// serviceAPI is the minimal subset of functions that we use from the AWS SDK, this allows for a greatly reduce surface
// area for mock generation.
type serviceAPI interface {
	AbortMultipartUpload(
		context.Context, *s3.AbortMultipartUploadInput, ...OptionFn,
	) (*s3.AbortMultipartUploadOutput, error)

	CompleteMultipartUpload(
		context.Context, *s3.CompleteMultipartUploadInput, ...OptionFn,
	) (*s3.CompleteMultipartUploadOutput, error)

	CreateMultipartUpload(
		context.Context, *s3.CreateMultipartUploadInput, ...OptionFn,
	) (*s3.CreateMultipartUploadOutput, error)

	DeleteObjects(context.Context, *s3.DeleteObjectsInput, ...OptionFn) (*s3.DeleteObjectsOutput, error)
	GetObject(context.Context, *s3.GetObjectInput, ...OptionFn) (*s3.GetObjectOutput, error)
	HeadObject(context.Context, *s3.HeadObjectInput, ...OptionFn) (*s3.HeadObjectOutput, error)

	ListObjectsV2Pages(
		context.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...OptionFn,
	) error

	ListPartsPages(
		context.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...OptionFn,
	) error

	PutObject(context.Context, *s3.PutObjectInput, ...OptionFn) (*s3.PutObjectOutput, error)
	UploadPart(context.Context, *s3.UploadPartInput, ...OptionFn) (*s3.UploadPartOutput, error)

	UploadPartCopy(
		context.Context, *s3.UploadPartCopyInput, ...OptionFn,
	) (*s3.UploadPartCopyOutput, error)
}
