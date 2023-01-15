## Huawei Cloud ObjectStorage Go SDK

This repository is a fork of [huaweicloud-sdk-go-obs](https://github.com/huaweicloud/huaweicloud-sdk-go-obs). Licensed under Apache 2.0.

### How to use

**Install**

```
go get github.com/logiqbits/go-huaweicloud-obs-sdk
```

**Create OBS Client**

```go

const (
  endpoint  = "https://obs.region.myhuaweicloud.com"
  accessKey = ""
  secretKey = ""
  bucket    = ""
  customDomain = "https://mydomain" // if want to alias with custom domain
)

client, err := obs.NewClient(endpoint, accessKey, secretKey)
if err != nil {
  panic(err)
}

client.WithAliasDomain(customDomain) // if want to alias with custom domain

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

See on `test` files for more demo.

Currently this package is only designed for LogiQbits internal purpose. All functions are not implemented yet, we'll update this package gradually as for our need. If someone want to contribute, we'll happy to receive pull request.
