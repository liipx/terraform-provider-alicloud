package dcdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDcdnOfflineLogDelivery invokes the dcdn.DescribeDcdnOfflineLogDelivery API synchronously
func (client *Client) DescribeDcdnOfflineLogDelivery(request *DescribeDcdnOfflineLogDeliveryRequest) (response *DescribeDcdnOfflineLogDeliveryResponse, err error) {
	response = CreateDescribeDcdnOfflineLogDeliveryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnOfflineLogDeliveryWithChan invokes the dcdn.DescribeDcdnOfflineLogDelivery API asynchronously
func (client *Client) DescribeDcdnOfflineLogDeliveryWithChan(request *DescribeDcdnOfflineLogDeliveryRequest) (<-chan *DescribeDcdnOfflineLogDeliveryResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnOfflineLogDeliveryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnOfflineLogDelivery(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDcdnOfflineLogDeliveryWithCallback invokes the dcdn.DescribeDcdnOfflineLogDelivery API asynchronously
func (client *Client) DescribeDcdnOfflineLogDeliveryWithCallback(request *DescribeDcdnOfflineLogDeliveryRequest, callback func(response *DescribeDcdnOfflineLogDeliveryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnOfflineLogDeliveryResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnOfflineLogDelivery(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDcdnOfflineLogDeliveryRequest is the request struct for api DescribeDcdnOfflineLogDelivery
type DescribeDcdnOfflineLogDeliveryRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnOfflineLogDeliveryResponse is the response struct for api DescribeDcdnOfflineLogDelivery
type DescribeDcdnOfflineLogDeliveryResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Fields    []string `json:"Fields" xml:"Fields"`
	Domains   []Domain `json:"Domains" xml:"Domains"`
	Regions   []Region `json:"Regions" xml:"Regions"`
}

// CreateDescribeDcdnOfflineLogDeliveryRequest creates a request to invoke DescribeDcdnOfflineLogDelivery API
func CreateDescribeDcdnOfflineLogDeliveryRequest() (request *DescribeDcdnOfflineLogDeliveryRequest) {
	request = &DescribeDcdnOfflineLogDeliveryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnOfflineLogDelivery", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnOfflineLogDeliveryResponse creates a response to parse from DescribeDcdnOfflineLogDelivery response
func CreateDescribeDcdnOfflineLogDeliveryResponse() (response *DescribeDcdnOfflineLogDeliveryResponse) {
	response = &DescribeDcdnOfflineLogDeliveryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
