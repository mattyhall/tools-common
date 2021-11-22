// Code generated by mockery v2.9.4. DO NOT EDIT.

package objazure

import (
	context "context"

	azblob "github.com/Azure/azure-storage-blob-go/azblob"

	io "io"

	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// mockBlockBlobAPI is an autogenerated mock type for the blockBlobAPI type
type mockBlockBlobAPI struct {
	mock.Mock
}

// CommitBlockList provides a mock function with given fields: ctx, base64BlockIDs, h, metadata, ac, tier, blobTagsMap, cpk
func (_m *mockBlockBlobAPI) CommitBlockList(ctx context.Context, base64BlockIDs []string, h azblob.BlobHTTPHeaders, metadata azblob.Metadata, ac azblob.BlobAccessConditions, tier azblob.AccessTierType, blobTagsMap azblob.BlobTagsMap, cpk azblob.ClientProvidedKeyOptions) (*azblob.BlockBlobCommitBlockListResponse, error) {
	ret := _m.Called(ctx, base64BlockIDs, h, metadata, ac, tier, blobTagsMap, cpk)

	var r0 *azblob.BlockBlobCommitBlockListResponse
	if rf, ok := ret.Get(0).(func(context.Context, []string, azblob.BlobHTTPHeaders, azblob.Metadata, azblob.BlobAccessConditions, azblob.AccessTierType, azblob.BlobTagsMap, azblob.ClientProvidedKeyOptions) *azblob.BlockBlobCommitBlockListResponse); ok {
		r0 = rf(ctx, base64BlockIDs, h, metadata, ac, tier, blobTagsMap, cpk)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*azblob.BlockBlobCommitBlockListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, azblob.BlobHTTPHeaders, azblob.Metadata, azblob.BlobAccessConditions, azblob.AccessTierType, azblob.BlobTagsMap, azblob.ClientProvidedKeyOptions) error); ok {
		r1 = rf(ctx, base64BlockIDs, h, metadata, ac, tier, blobTagsMap, cpk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StageBlock provides a mock function with given fields: ctx, base64BlockID, body, ac, transactionalMD5, cpk
func (_m *mockBlockBlobAPI) StageBlock(ctx context.Context, base64BlockID string, body io.ReadSeeker, ac azblob.LeaseAccessConditions, transactionalMD5 []byte, cpk azblob.ClientProvidedKeyOptions) (*azblob.BlockBlobStageBlockResponse, error) {
	ret := _m.Called(ctx, base64BlockID, body, ac, transactionalMD5, cpk)

	var r0 *azblob.BlockBlobStageBlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, io.ReadSeeker, azblob.LeaseAccessConditions, []byte, azblob.ClientProvidedKeyOptions) *azblob.BlockBlobStageBlockResponse); ok {
		r0 = rf(ctx, base64BlockID, body, ac, transactionalMD5, cpk)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*azblob.BlockBlobStageBlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, io.ReadSeeker, azblob.LeaseAccessConditions, []byte, azblob.ClientProvidedKeyOptions) error); ok {
		r1 = rf(ctx, base64BlockID, body, ac, transactionalMD5, cpk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StageBlockFromURL provides a mock function with given fields: ctx, base64BlockID, sourceURL, offset, count, destinationAccessConditions, sourceAccessConditions, cpk
func (_m *mockBlockBlobAPI) StageBlockFromURL(ctx context.Context, base64BlockID string, sourceURL url.URL, offset, count int64, destinationAccessConditions azblob.LeaseAccessConditions, sourceAccessConditions azblob.ModifiedAccessConditions, cpk azblob.ClientProvidedKeyOptions) (*azblob.BlockBlobStageBlockFromURLResponse, error) {
	ret := _m.Called(ctx, base64BlockID, sourceURL, offset, count, destinationAccessConditions, sourceAccessConditions, cpk)

	var r0 *azblob.BlockBlobStageBlockFromURLResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, url.URL, int64, int64, azblob.LeaseAccessConditions, azblob.ModifiedAccessConditions, azblob.ClientProvidedKeyOptions) *azblob.BlockBlobStageBlockFromURLResponse); ok {
		r0 = rf(ctx, base64BlockID, sourceURL, offset, count, destinationAccessConditions, sourceAccessConditions, cpk)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*azblob.BlockBlobStageBlockFromURLResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, url.URL, int64, int64, azblob.LeaseAccessConditions, azblob.ModifiedAccessConditions, azblob.ClientProvidedKeyOptions) error); ok {
		r1 = rf(ctx, base64BlockID, sourceURL, offset, count, destinationAccessConditions, sourceAccessConditions, cpk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URL provides a mock function with given fields:
func (_m *mockBlockBlobAPI) URL() url.URL {
	ret := _m.Called()

	var r0 url.URL
	if rf, ok := ret.Get(0).(func() url.URL); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(url.URL)
	}

	return r0
}

// Upload provides a mock function with given fields: ctx, body, h, metadata, ac, tier, blobTagsMap, cpk
func (_m *mockBlockBlobAPI) Upload(ctx context.Context, body io.ReadSeeker, h azblob.BlobHTTPHeaders, metadata azblob.Metadata, ac azblob.BlobAccessConditions, tier azblob.AccessTierType, blobTagsMap azblob.BlobTagsMap, cpk azblob.ClientProvidedKeyOptions) (*azblob.BlockBlobUploadResponse, error) {
	ret := _m.Called(ctx, body, h, metadata, ac, tier, blobTagsMap, cpk)

	var r0 *azblob.BlockBlobUploadResponse
	if rf, ok := ret.Get(0).(func(context.Context, io.ReadSeeker, azblob.BlobHTTPHeaders, azblob.Metadata, azblob.BlobAccessConditions, azblob.AccessTierType, azblob.BlobTagsMap, azblob.ClientProvidedKeyOptions) *azblob.BlockBlobUploadResponse); ok {
		r0 = rf(ctx, body, h, metadata, ac, tier, blobTagsMap, cpk)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*azblob.BlockBlobUploadResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, io.ReadSeeker, azblob.BlobHTTPHeaders, azblob.Metadata, azblob.BlobAccessConditions, azblob.AccessTierType, azblob.BlobTagsMap, azblob.ClientProvidedKeyOptions) error); ok {
		r1 = rf(ctx, body, h, metadata, ac, tier, blobTagsMap, cpk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}