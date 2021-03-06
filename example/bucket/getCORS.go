package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"net/http"

	"github.com/huanght1997/cos-go-sdk-v5"
	"github.com/huanght1997/cos-go-sdk-v5/debug"
)

func main() {
	u, _ := url.Parse("https://test-1253846586.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{
		BucketURL: u,
	}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	v, _, err := c.Bucket.GetCORS(context.Background())
	if err != nil {
		panic(err)
	}
	for _, r := range v.Rules {

		fmt.Printf("%s, %s\n", r.AllowedOrigins, r.AllowedMethods)
	}
}
