package DescribeDcdnDomainCertificateInfo

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
)

func DescribeDcdnDomainCertificateInfo(AccessKeyID string, AccessKeySecret string, DomainName string) string {
	client, err := dcdn.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := dcdn.CreateDescribeDcdnDomainCertificateInfoRequest()
	request.Scheme = "https"

	request.DomainName = DomainName

	response, err := client.DescribeDcdnDomainCertificateInfo(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return response.String()
}
