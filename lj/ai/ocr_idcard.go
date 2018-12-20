package main

import (
	"fmt"
	"os"

	"github.com/qiniu/api.v7/auth/qbox"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
)

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
	domain    = os.Getenv("QINIU_TEST_DOMAIN")
	bucket    = os.Getenv("QINIU_TEST_BUCKET")
)

type OcrIdcard struct {
	Data *OcrIdcardData `json:"data"`
}

type OcrIdcardData struct {
	Uri string `json:"uri"`
}

func main() {
	mac := qbox.NewMac(accessKey, secretKey)

	url := "http://ai.qiniuapi.com/v1/ocr/idcard"
	method := "POST"
	host := "ai.qiniuapi.com"
	contentType := "application/json"
	bodyUri := OcrIdcardData{Uri: "http://test-pub.iamlj.com/test-idcard.jpg"}
	body := OcrIdcard{Data: &bodyUri}

	reqData, _ := json.Marshal(body)

	req, reqErr := http.NewRequest(method, url, bytes.NewReader(reqData))
	if reqErr != nil {
		return
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Host", host)

	qiniuToken, signErr := mac.SignRequestV2(req)
	if signErr != nil {
		fmt.Printf(signErr.Error())
	}

	req.Header.Add("Authorization", "Qiniu "+qiniuToken)

	fmt.Println(string(url))
	fmt.Println(string(reqData))
	fmt.Println(string("Qiniu " + qiniuToken))

	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		fmt.Printf(respErr.Error())
	}
	defer resp.Body.Close()

	resData, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		fmt.Printf(ioErr.Error())
	}

	fmt.Println(string(resData))
}
