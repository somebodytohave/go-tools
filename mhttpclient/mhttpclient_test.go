package mhttpclient_test

import (
	"fmt"
	"github.com/sun-wenming/go-tools/mhttpclient"
	"net/url"
	"testing"
)

func TestHttpClient(t *testing.T) {
	var urls = ""
	var req = mhttpclient.NewRequestPost(urls)
	var results interface{}
	var params = url.Values{}
	params.Add("key", "value")
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.SetParams(params)
	var rep = req.Exec()
	err := rep.Unmarshal(&results)
	fmt.Println(results)
	//fmt.Println(rep.String())
	t.Log(err)
}
