package obs_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/logiqbits/go-huaweicloud-obs-sdk/obs"
)

const (
	endpoint     = "https://obs.ap-southeast-3.myhuaweicloud.com"
	accessKey    = "JBH8RXIAAE9X9HBWKANB"
	secretKey    = "8W1hSERaKD6aUTiqXEq96N7woKLXRD78c41bJeeH"
	bucket       = "logiqbits-public-bucket"
	customDomain = "https://static-cdn.shopap.io"
)

func TestSimpleFileUpload(t *testing.T) {
	client, err := obs.NewClient(endpoint, accessKey, secretKey)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	client.WithAliasDomain(customDomain)

	sourceFile, _ := os.Open("/Users/rafi/Desktop/rafiul-islam.pdf")
	body := bufio.NewReader(sourceFile)

	res, err := client.UploadObjectToBucket(bucket, "business", "cv.pdf", body, nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(res)
}

func TestFileDelete(t *testing.T) {
	client, err := obs.NewClient(endpoint, accessKey, secretKey)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	res, err := client.DeleteObjectFromBucket(bucket, "1673702033250953000.pdf")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(res)

}

func TestDeleteFolderFromBucket(t *testing.T) {
	client, err := obs.NewClient(endpoint, accessKey, secretKey)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	res, err := client.DeleteFolderFromBucket(bucket, "personal")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(res)
}

func TestDeleteMultipleFilesFromBucket(t *testing.T) {
	files := []string{
		"rafiul/cv.pdf",
		"1673701374376324000.pdf",
	}

	client, err := obs.NewClient(endpoint, accessKey, secretKey)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	res, err := client.DeleteObjectsFromBucket(bucket, files)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(res)
}
