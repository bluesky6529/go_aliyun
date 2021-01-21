package RefreshObjectCaches

import (
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
)

func RefreshObjectCaches(AccessKeyID string, AccessKeySecret string, ObjectPath []string, dir string) string {
	client, err := cdn.NewClientWithAccessKey("cn-hanzhou", AccessKeyID, AccessKeySecret)

	request := cdn.CreateRefreshObjectCachesRequest()
	request.Scheme = "https"

	request.ObjectPath = strings.Join(ObjectPath, "\n")
	request.ObjectType = dir

	response, err := client.RefreshObjectCaches(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	return response.String()
}
