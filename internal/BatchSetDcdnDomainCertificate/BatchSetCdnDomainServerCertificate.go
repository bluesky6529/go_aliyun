package BatchSetDcdnDomainCertificate

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
)

func BatchSetDcdnDomainCertificate(AccessKeyID string, AccessKeySecret string, DomainName string, SSLPub string, SSLPri string) string {
	client, err := dcdn.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := dcdn.CreateBatchSetDcdnDomainCertificateRequest()
	request.Scheme = "https"

	request.CertType = "upload"
	request.SSLProtocol = "on"
	request.DomainName = DomainName
	request.SSLPub = SSLPub
	request.SSLPri = SSLPri

	response, err := client.BatchSetDcdnDomainCertificate(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return response.String()
}
