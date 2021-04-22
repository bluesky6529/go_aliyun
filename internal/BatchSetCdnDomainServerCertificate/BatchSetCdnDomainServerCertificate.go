package BatchSetCdnDomainServerCertificate

import (
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
)

func BatchSetCdnDomainServerCertificate(AccessKeyID string, AccessKeySecret string, DomainName string, SSLPub string, SSLPri string) string {
	client, err := cdn.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := cdn.CreateBatchSetCdnDomainServerCertificateRequest()
	request.Scheme = "https"

	request.DomainName = DomainName
	request.SSLProtocol = "on"
	request.CertName = DomainName + "--" + time.Now().Format("20060102")
	request.CertType = "upload"
	request.SSLPub = SSLPub
	request.SSLPri = SSLPri

	response, err := client.BatchSetCdnDomainServerCertificate(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return request.SSLPri + "\n" + request.SSLPub
}
