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
	u, _ := url.Parse("https://alanbj-1251668577.cos.ap-beijing.myqcloud.com")
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

	v, _, err := c.Bucket.GetBucketReplication(context.Background())
	if err != nil {
		panic(err)
	}
	for _, r := range v.Rule {

		fmt.Printf("%s, %s\n", r.Status, r.ID)
	}
}
