package obs_test

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/logiqbits/go-huaweicloud-obs-sdk/obs"
)

const (
	endpoint  = "https://myhuaweicloud.com"
	accessKey = ""
	secretKey = ""
	bucket    = ""
)

func TestSimpleFileUpload(t *testing.T) {
	client, err := obs.NewClient(endpoint, accessKey, secretKey)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	sourceFile, _ := os.Open("/Users/rafi/Desktop/rafiul-islam.pdf")
	body := bufio.NewReader(sourceFile)

	res, err := client.UploadFileToBucket(bucket, "rafiul", "cv.pdf", body, nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(res)
	log.Println(res)
}
