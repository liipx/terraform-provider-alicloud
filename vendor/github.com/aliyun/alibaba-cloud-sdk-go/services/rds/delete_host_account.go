package rds

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

// DeleteHostAccount invokes the rds.DeleteHostAccount API synchronously
// api document: https://help.aliyun.com/api/rds/deletehostaccount.html
func (client *Client) DeleteHostAccount(request *DeleteHostAccountRequest) (response *DeleteHostAccountResponse, err error) {
	response = CreateDeleteHostAccountResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHostAccountWithChan invokes the rds.DeleteHostAccount API asynchronously
// api document: https://help.aliyun.com/api/rds/deletehostaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHostAccountWithChan(request *DeleteHostAccountRequest) (<-chan *DeleteHostAccountResponse, <-chan error) {
	responseChan := make(chan *DeleteHostAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHostAccount(request)
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

// DeleteHostAccountWithCallback invokes the rds.DeleteHostAccount API asynchronously
// api document: https://help.aliyun.com/api/rds/deletehostaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHostAccountWithCallback(request *DeleteHostAccountRequest, callback func(response *DeleteHostAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHostAccountResponse
		var err error
		defer close(result)
		response, err = client.DeleteHostAccount(request)
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

// DeleteHostAccountRequest is the request struct for api DeleteHostAccount
type DeleteHostAccountRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DeleteHostAccountResponse is the response struct for api DeleteHostAccount
type DeleteHostAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteHostAccountRequest creates a request to invoke DeleteHostAccount API
func CreateDeleteHostAccountRequest() (request *DeleteHostAccountRequest) {
	request = &DeleteHostAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteHostAccount", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteHostAccountResponse creates a response to parse from DeleteHostAccount response
func CreateDeleteHostAccountResponse() (response *DeleteHostAccountResponse) {
	response = &DeleteHostAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}