package QueryAccountBalance

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func QueryAccountBalance(AccessKeyID string, AccessKeySecret string) string {
	client, err := bssopenapi.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := bssopenapi.CreateQueryAccountBalanceRequest()
	request.Scheme = "https"

	response, err := client.QueryAccountBalance(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Printf("response is %#v\n", response.Data.AvailableAmount)
	return response.Data.AvailableAmount
}
