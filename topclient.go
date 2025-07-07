package topsdk

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
	"topsdk/util"
)

type TopClient struct {
	AppKey string
	AppSecret string
	ServerUrl string
	Format string
	SignMethod string
	ConnectTimeout int64
	ReadTimeout int64
	Version string
	Simplify bool
	httpClient *http.Client
}
type HttpTransportConfig struct {
	DialTimeout int64
	KeepAlive int64
	MaxIdleConns int
	MaxIdleConnsPerHost int
	IdleConnTimeout int64
	MaxConnsPerHost int

}

func NewDefaultTopClient(AppKey string,AppSecret string,ServerUrl string, connectTimeount int64, readTimeout int64) TopClient {
	var httpTransportConfig = &HttpTransportConfig{
		DialTimeout: 30000,
		KeepAlive: 30000,
		MaxIdleConns: 100,
		MaxIdleConnsPerHost: 50,
		IdleConnTimeout: 30000,
	}
	return NewTopClientWithConfig(AppKey,AppSecret,ServerUrl,connectTimeount,readTimeout,httpTransportConfig)

}


func NewTopClientWithConfig(AppKey string,AppSecret string,ServerUrl string, connectTimeount int64, readTimeout int64,httpTransportConfig *HttpTransportConfig) TopClient {
	httpClient := http.Client{
		Timeout: time.Duration(connectTimeount) * time.Millisecond,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Duration(httpTransportConfig.DialTimeout) * time.Millisecond,
				KeepAlive: time.Duration(httpTransportConfig.KeepAlive) * time.Millisecond,
			}).DialContext,
			ForceAttemptHTTP2: true,
			MaxIdleConns: httpTransportConfig.MaxIdleConns,
			MaxIdleConnsPerHost: httpTransportConfig.MaxIdleConnsPerHost,
			IdleConnTimeout: time.Duration(httpTransportConfig.IdleConnTimeout) * time.Millisecond,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	return TopClient{
		AppKey:         AppKey,
		AppSecret:      AppSecret,
		ServerUrl:      ServerUrl,
		Format:         ApiFormat,
		SignMethod:     SignMethod,
		ConnectTimeout: connectTimeount,
		ReadTimeout:    readTimeout,
		Version:        TopVersion,
		Simplify:       true,
		httpClient:     &httpClient,
	}
}




func (client *TopClient) ExecuteWithSession(method string,data map[string]interface{},fileData map[string]interface{},session string) (string,error){
	var publicParam = make(map[string]interface{})
	publicParam["method"] = method
	publicParam["app_key"] = client.AppKey
	publicParam["timestamp"] = time.Now().Format(DateFormat)
	publicParam["v"] = client.Version
	publicParam["sign_method"] = client.SignMethod
	publicParam["format"] = client.Format
	publicParam["simplify"] = client.Simplify
	publicParam["partner_id"] = SdkVersion
	if session != "" {
		publicParam["session"] = session
	}
	sign := util.GetSign(publicParam, data, client.AppSecret)
	// 构建url
	serverUrl, _ := url.Parse(client.ServerUrl)
	urlValues := url.Values{}
	urlValues.Add("sign",sign)
	for k,v := range publicParam{
		urlValues.Add(k,fmt.Sprint(v))
	}
	serverUrl.RawQuery = urlValues.Encode()
	urlPath := serverUrl.String()
	// 构建body
	if fileData != nil && len(fileData) > 0 {
		return doPostWithFile(urlPath,data,fileData,client.httpClient)
	}else {
		return doPost(urlPath, data, client.httpClient)
	}

}

func doPost(urlPath string,data map[string]interface{},httpClient *http.Client) (string,error){
	bodyParam := url.Values{}
	for k,v := range data{
		bodyParam.Add(k,fmt.Sprint(v))
	}
	resp, err := httpClient.Post(urlPath,"application/x-www-form-urlencoded",strings.NewReader(bodyParam.Encode()))
	if resp != nil {
		defer resp.Body.Close()
	}
	if(err != nil){
		log.Println("http.PostForm error",err)
		return "",err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if(err != nil){
		log.Println("ioutil.ReadAll",err)
		return "",err
	}
	return string(body),nil
}

func doPostWithFile(urlPath string,data map[string]interface{},fileData map[string]interface{}, httpClient *http.Client) (string,error){
	bodyBuf := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuf)
	for k,v := range data{
		err := writer.WriteField(k, fmt.Sprint(v))
		if err != nil {
			return "", err
		}
	}
	for k,v := range fileData {
		value , ok := v.([]byte)
		if ok {
			fileWriter, err := writer.CreateFormFile(k, "file")
			if err != nil {
				return "",err
			}
			_, err = io.Copy(fileWriter, bytes.NewReader(value))
			if err != nil {
				return "",err
			}
		}else {
			value , ok := v.(*util.FileItem)
			if ok {
				fileWriter, err := writer.CreateFormFile(k, value.FileName)
				if err != nil {
					return "",err
				}
				_, err = io.Copy(fileWriter, bytes.NewReader(value.Content))
				if err != nil {
					return "",err
				}
			}
		}
	}

	err := writer.Close()
	if(err != nil){
		return "",err
	}

	resp, err := httpClient.Post(urlPath,writer.FormDataContentType(),bodyBuf)
	if err != nil {
		log.Println("http.PostForm error",err)
		return "",err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll",err)
		return "",err
	}
	return string(body),nil
}


func (client *TopClient) Execute(method string,data map[string]interface{},fileData map[string]interface{}) (string,error){
	return client.ExecuteWithSession(method,data,fileData,"")
}



