package linuxkit

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"encoding/json"
	"strings"
)

const serverIP = "10.10.98.2"
const serverPort = "5000"

// ASBDDSClient contains state required for communication with ASBDDS
type ASBDDSClient struct {
	serverIP string
	serverPort string
	apiBaseURL string
	rest resty.Client
}

type ASBDDSResponseStatus struct {
	Code int
	Message string
}

type ASBDDSResponse struct {
	Status ASBDDSResponseStatus
	Data json.RawMessage
}

// NewASBDDSClient creates a new ASBDDS client
func NewASBDDSClient() (*ASBDDSClient, error) {
	var client = &ASBDDSClient{
		serverIP: serverIP,
		serverPort: serverPort,
		rest: *resty.New(),
		apiBaseURL: "http://" + serverIP + ":" + serverPort + "/",
	}
	return client, nil
}

func CheckResponse(resp resty.Response) (*json.RawMessage, error) {
	var response ASBDDSResponse
	err := json.Unmarshal([]byte(resp.String()), &response)
	if err != nil {
		err = fmt.Errorf("unable to parse response from asbdds api, please check server access")
	}

	if response.Status.Code != 0 {
		if len(response.Status.Message) > 3 {
			err = fmt.Errorf(strings.ToLower(response.Status.Message))
		} else {
			code := response.Status.Code
			if code == 1 {
				err = fmt.Errorf("requested object not found")
			} else if code == 2 {
				err = fmt.Errorf("error, bad request")
			} else if code == 3 {
				err = fmt.Errorf("try later")
			}
		}
	}
	return &response.Data, err
}

// CreateDevice create a device in asbdds
func (a ASBDDSClient) CreateDevice(model, ipxeURL string) (*json.RawMessage, error){
	resp, err := a.rest.R().
		SetQueryParams(map[string]string{
			"model": model,
			"ipxe_url": ipxeURL,
		}).
		SetHeader("Accept", "application/json").
		Put(a.apiBaseURL + "device")

	if err != nil {
		return nil, err
	}

	return CheckResponse(*resp)
}

// DeleteDevice delete a device in asbdds
func (a ASBDDSClient) DeleteDevice(deviceUUID string) (*json.RawMessage, error){
	resp, err := a.rest.R().
		SetHeader("Accept", "application/json").
		Delete(a.apiBaseURL + "device/" + deviceUUID)

	if err != nil {
		return nil, err
	}

	return CheckResponse(*resp)
}