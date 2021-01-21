package RefreshDcdnObjectCaches

import (
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
)

func RefreshDcdnObjectCaches(AccessKeyID string, AccessKeySecret string, ObjectPath []string, dir string) string {
	client, err := dcdn.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := dcdn.CreateRefreshDcdnObjectCachesRequest()
	request.Scheme = "https"

	request.ObjectPath = strings.Join(ObjectPath, "\n")
	request.ObjectType = dir

	response, err := client.RefreshDcdnObjectCaches(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return response.String()
}
