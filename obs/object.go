package obs

import (
	"fmt"
	"io"
	"net/url"
	"path"

	"github.com/logiqbits/go-huaweicloud-obs-sdk/internal/sdk"
)

type SingleUploadResponse struct {
	Location  string
	ETag      string
	VersionId string
	Headers   map[string][]string
}

func (c *Client) UploadObjectToBucket(bucket, folder, fileName string, body io.Reader, metadata map[string]string) (*SingleUploadResponse, error) {
	objectKey := fileName
	if folder != "" {
		objectKey = path.Join(folder, fileName)
	}
	input := &sdk.PutObjectInput{}
	input.Bucket = bucket
	input.Key = objectKey
	input.Metadata = metadata
	input.Body = body
	output, err := c.obsClient.PutObject(input)
	if err != nil {
		return nil, err
	}

	var location string
	if c.aliasDomain != "" {
		location = fmt.Sprintf("%s/%s", c.aliasDomain, objectKey)
	} else {
		hostUrl, _ := url.Parse(c.endpoint)
		location = fmt.Sprintf("%s://%s.%s/%s", hostUrl.Scheme, bucket, hostUrl.Host, objectKey)
	}

	res := &SingleUploadResponse{
		ETag:      output.ETag,
		VersionId: output.VersionId,
		Location:  location,
		Headers:   output.ResponseHeaders,
	}

	return res, nil
}

type ObjectDeleteResponse struct {
	VersionId string
	Headers   map[string][]string
}

func (c *Client) DeleteObjectFromBucket(bucket, filePath string) (*ObjectDeleteResponse, error) {
	input := &sdk.DeleteObjectInput{}
	input.Bucket = bucket
	input.Key = filePath
	output, err := c.obsClient.DeleteObject(input)

	if err != nil {
		return nil, err

	}

	res := &ObjectDeleteResponse{VersionId: output.VersionId, Headers: output.ResponseHeaders}
	return res, nil
}

type ObjectsDeleteResponse struct {
	DeleteItems []sdk.Deleted
	Headers     map[string][]string
}

func (c *Client) DeleteFolderFromBucket(bucket, folderPath string) (*ObjectsDeleteResponse, error) {
	listInput := &sdk.ListObjectsInput{}
	listInput.Bucket = bucket
	listInput.Prefix = folderPath
	listOutput, err := c.obsClient.ListObjects(listInput)
	if err != nil {
		return nil, err
	}

	if len(listOutput.Contents) == 0 {
		return &ObjectsDeleteResponse{DeleteItems: make([]sdk.Deleted, 0), Headers: map[string][]string{}}, nil
	}

	bulkDeleteInput := sdk.DeleteObjectsInput{}
	bulkDeleteInput.Bucket = bucket
	bulkDeleteInput.Objects = make([]sdk.ObjectToDelete, 0)
	for _, c := range listOutput.Contents {
		bulkDeleteInput.Objects = append(bulkDeleteInput.Objects, sdk.ObjectToDelete{
			Key: c.Key,
		})
	}

	bulkDeleteOutput, err := c.obsClient.DeleteObjects(&bulkDeleteInput)
	if err != nil {
		return nil, err
	}

	res := &ObjectsDeleteResponse{
		DeleteItems: bulkDeleteOutput.Deleteds,
		Headers:     bulkDeleteOutput.ResponseHeaders,
	}

	return res, nil
}

func (c *Client) DeleteObjectsFromBucket(bucket string, filePaths []string) (*ObjectsDeleteResponse, error) {
	bulkDeleteInput := sdk.DeleteObjectsInput{}
	bulkDeleteInput.Bucket = bucket
	bulkDeleteInput.Objects = make([]sdk.ObjectToDelete, 0)
	for _, path := range filePaths {
		bulkDeleteInput.Objects = append(bulkDeleteInput.Objects, sdk.ObjectToDelete{
			Key: path,
		})
	}

	bulkDeleteOutput, err := c.obsClient.DeleteObjects(&bulkDeleteInput)
	if err != nil {
		return nil, err
	}

	res := &ObjectsDeleteResponse{
		DeleteItems: bulkDeleteOutput.Deleteds,
		Headers:     bulkDeleteOutput.ResponseHeaders,
	}

	return res, nil
}
