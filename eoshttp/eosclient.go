package eoshttp

import (
	"net/http"
	"bytes"
	"fmt"
	"log"
	"encoding/base64"
	"io/ioutil"
	"encoding/json"
)

const (
	Version              = "v1"
	EOSLocalhostEndpoint = "http://localhost:8888"

	RequestTypeJSON = "json"
	RequestTypeText = "texts"
)

type EOSClient struct {
	Version       string
	endpoint      string
	httpClient    *http.Client
	customHeaders map[string]string
	Chain         *Chain
	Wallet        *Wallet
	Logging       bool
}

type EOSRequest struct {
	Path    string
	SubPath string
	Body    []byte
	Type    string
}

type EOSResponse struct {
	Status     string
	StatusCode int
	Success    bool
	Error      error
	EOSError   *Error
	IsEOSError bool
	Body       []byte
}

func NewEOSClient() *EOSClient {
	client := &EOSClient{
		Version:    Version,
		endpoint:   createURL(EOSLocalhostEndpoint, Version),
		httpClient: http.DefaultClient,
		Logging:    true,
	}

	client.Chain = newChain(client)
	client.Wallet = newWallet(client)

	return client
}

func (client *EOSClient) WithEndpoint(endpoint string) *EOSClient {
	client.endpoint = createURL(EOSLocalhostEndpoint, Version)
	return client
}

func (client *EOSClient) WithCustomHeader(key string, value string) *EOSClient {
	client.customHeaders[key] = value
	return client
}

func (client *EOSClient) UnsetCustomHeader(key string) *EOSClient {
	delete(client.customHeaders, key)
	return client
}

func (client *EOSClient) SetBasicAuth(username string, password string) *EOSClient {
	if username == "" || password == "" {
		delete(client.customHeaders, "Authorization")
	} else {
		auth := username + ":" + password
		client.customHeaders["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	}

	return client
}

// SetHTTPClient can be used to set a custom http.Client.
// This can be useful for example if you want to customize the http.Client behaviour (e.g. proxy settings)
func (client *EOSClient) WithHTTPClient(httpClient *http.Client) *EOSClient {
	if httpClient == nil {
		panic("httpClient cannot be nil")
	}

	client.httpClient = httpClient

	return client
}

func (client *EOSClient) NewHttpRequest(request *EOSRequest) (*http.Request, error) {
	url := fmt.Sprintf("%v/%v/%v", client.endpoint, request.Path, request.SubPath)

	httpRequest, err := http.NewRequest("POST", url, bytes.NewReader(request.Body))

	if err != nil {
		return nil, err
	}

	for k, v := range client.customHeaders {
		httpRequest.Header.Add(k, v)
	}

	switch request.Type {
	case RequestTypeJSON:
		httpRequest.Header.Add("Content-Type", "application/json")
		httpRequest.Header.Add("Accept", "application/json")
	case RequestTypeText:
		httpRequest.Header.Add("Content-Type", "application/text")
		httpRequest.Header.Add("Accept", "application/text")
	}

	log.Printf("Calling %v with Body %+v", url, string(request.Body))

	return httpRequest, nil
}

func (client *EOSClient) Call(request *EOSRequest) *EOSResponse {
	httpRequest, err := client.NewHttpRequest(request)

	if err != nil {
		return client.NewErrorResponse(err, -1, "")
	}

	httpResponse, err := client.httpClient.Do(httpRequest)
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		return client.NewErrorResponse(err, httpResponse.StatusCode, httpResponse.Status)
	}

	if httpResponse.StatusCode == 500 {
		eoserror := Error{}
		json.Unmarshal(body, &eoserror)
		return client.NewEOSErrorResponse(&eoserror, httpResponse.StatusCode, httpResponse.Status)
	}

	if httpResponse.StatusCode >= 300 {
		return client.NewErrorResponse(fmt.Errorf("request returned with http error code: %v", httpResponse.Status), httpResponse.StatusCode, httpResponse.Status)
	}

	return &EOSResponse{
		Body:       body,
		Success:    true,
		Status:     httpResponse.Status,
		StatusCode: httpResponse.StatusCode,
	}
}

func (client *EOSClient) NewEOSJSONRequest(path string, subPath string, body []byte) *EOSRequest {
	return &EOSRequest{
		Path:    path,
		Body:    body,
		SubPath: subPath,
		Type:    RequestTypeJSON,
	}
}
func (client *EOSClient) NewEOSTextRequest(path string, subPath string, body []byte) *EOSRequest {
	return &EOSRequest{
		Path:    path,
		Body:    body,
		SubPath: subPath,
		Type:    RequestTypeText,
	}
}

func (client *EOSClient) NewErrorResponse(err error, statusCode int, status string) *EOSResponse {
	return &EOSResponse{
		Error:      err,
		Success:    false,
		StatusCode: statusCode,
		Status:     status,
	}
}

func (client *EOSClient) NewEOSErrorResponse(eosError *Error, statusCode int, status string) *EOSResponse {
	return &EOSResponse{
		EOSError:   eosError,
		IsEOSError: true,
		Success:    false,
		StatusCode: statusCode,
		Status:     status,
	}
}

func (client *EOSClient) Log(message string) {
	if client.Logging {
		log.Println(message)
	}
}

func createURL(endpint string, version string) string {
	return fmt.Sprintf("%v/%v", endpint, version)
}

/*
	Default Http Error
	Example
	{
  		"code": 500,
  		"message": "Internal Service Error",
  		"error": {
    		"code": 3120001,
    		"name": "wallet_exist_exception",
    		"what": "Wallet already exists",
    		"details": [{
        		"message": "Wallet with name: 'test1' already exists at /mnt/dev/data/./test1.wallet",
        		"file": "wallet_manager.cpp",
        		"line_number": 42,
        		"method": "create"
      		}]
  		}
	}
 */
type Error struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Error   ErrorInternal `json:"error"`
}

type ErrorInternal struct {
	Code    int64         `json:"code"`
	Name    string        `json:"name"`
	What    string        `json:"what"`
	Details []ErrorDetail `json:"details"`
}

type ErrorDetail struct {
	Message    string `json:"message"`
	File       string `json:"file"`
	LineNumber int    `json:"line_number"`
	Method     string `json:"method"`
}
