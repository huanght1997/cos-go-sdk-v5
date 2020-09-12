package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"fmt"
	"github.com/huanght1997/cos-go-sdk-v5"
	"github.com/huanght1997/cos-go-sdk-v5/debug"
)

func main() {
	uin := "100010805041"
	appid := 1259654469
	jobid := "795ad997-5557-4869-9a19-b66ec087d460"
	u, _ := url.Parse("https://" + uin + ".cos-control.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BatchURL: u}
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

	headers := &cos.BatchRequestHeaders{
		XCosAppid: appid,
	}

	res, _, err := c.Batch.DescribeJob(context.Background(), jobid, headers)
	if err != nil {
		panic(err)
	}
	if res != nil && res.Job != nil {
		fmt.Printf("%+v", res.Job)
	}
}
