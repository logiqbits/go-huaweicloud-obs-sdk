##  Huawei Cloud ObjectStorage Go SDK

This repository is a fork of [huaweicloud-sdk-go-obs](https://github.com/huaweicloud/huaweicloud-sdk-go-obs). Licensed under Apache 2.0.

### How to use

**Create OBS Client**

```go

const (
	endpoint  = "https://myhuaweicloud.com"
	accessKey = ""
	secretKey = ""
	bucket    = ""
)

client, err := obs.NewClient(endpoint, accessKey, secretKey)
if err != nil {
  panic(err)
}

```



**Upload a simple file**

```go
sourceFile, _ := os.Open("/Users/rafi/Desktop/rafiul-islam.pdf")
body := bufio.NewReader(sourceFile)
res, err := client.UploadFileToBucket(bucket, "rafiul", "cv.pdf", body, nil)
if err != nil {
  panic(err)
}
log.Println(res)
```

