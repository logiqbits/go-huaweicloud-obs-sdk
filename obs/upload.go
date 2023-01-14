package obs

import (
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"

	"github.com/logiqbits/go-huaweicloud-obs-sdk/internal/sdk"
)

type UploadResponse struct {
	Location  string
	ETag      string
	VersionId string
}

type Client struct {
	endpoint  string
	obsClient *sdk.ObsClient
}

func NewClient(endpoint, accessKey, secretKey string) (*Client, error) {
	obsClient, err := sdk.New(accessKey, secretKey, endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{obsClient: obsClient, endpoint: endpoint}, nil
}

func (c *Client) UploadFileToBucket(bucket, folder, fileName string, body io.Reader, metadata map[string]string) (*UploadResponse, error) {
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

	hostUrl, _ := url.Parse(c.endpoint)
	var location string
	// checking whether user is using default endpoint or not
	if strings.Contains(hostUrl.Host, "myhuaweicloud.com") {
		// default endpoint
		location = fmt.Sprintf("%s://%s.%s/%s", hostUrl.Scheme, bucket, hostUrl.Host, objectKey)
	} else {
		// user custom domain endpoint
		location = fmt.Sprintf("%s://%s/%s", hostUrl.Scheme, hostUrl.Host, objectKey)
	}

	res := &UploadResponse{
		ETag:      output.ETag,
		VersionId: output.VersionId,
		Location:  location,
	}

	return res, nil
}
