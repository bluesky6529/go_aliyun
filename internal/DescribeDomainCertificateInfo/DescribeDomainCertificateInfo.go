package DescribeDomainCertificateInfo

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
)

func DescribeDomainCertificateInfo(AccessKeyID string, AccessKeySecret string, DomainName string) string {
	client, err := cdn.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := cdn.CreateDescribeDomainCertificateInfoRequest()
	request.Scheme = "https"

	request.DomainName = DomainName

	response, err := client.DescribeDomainCertificateInfo(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Printf("response is %#v\n", response)
	//return gjson.Get(response.CertInfos.CertInfo, "CertExpireTime").String()
	//const json = response.CertInfos.CertInfo
	return response.String()
}
