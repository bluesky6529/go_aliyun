package DescribeDcdnRefreshTasks

import (
	"fmt"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
)

func DescribeDcdnRefreshTasks(AccessKeyID string, AccessKeySecret string, url string) string {
	client, err := dcdn.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := dcdn.CreateDescribeDcdnRefreshTasksRequest()
	request.Scheme = "https"

	request.ObjectPath = url

	response, err := client.DescribeDcdnRefreshTasks(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Printf("response is %#v\n", response)

	//return response.String()

	var processInfo []string
	loc, _ := time.LoadLocation("Asia/Taipei")
	for _, t := range response.Tasks.Task {
		ct, _ := time.Parse(time.RFC3339, t.CreationTime)
		processInfo = append(processInfo, fmt.Sprintf("%s [%s] %s 進度:%s", ct.In(loc).Format("2006/01/02 15:04:05"), t.ObjectType, t.ObjectPath, t.Process))
	}
	return strings.Join(processInfo, "\n")
}
