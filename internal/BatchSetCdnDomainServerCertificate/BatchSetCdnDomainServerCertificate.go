package BatchSetCdnDomainServerCertificate

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
)

func BatchSetCdnDomainServerCertificate(AccessKeyID string, AccessKeySecret string) string {
	client, err := cdn.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)

	request := cdn.CreateBatchSetCdnDomainServerCertificateRequest()
	request.Scheme = "https"

	request.DomainName = "8j24k.cn"
	request.SSLProtocol = "on"
	request.CertName = "8j24k.cn"
	request.CertType = "upload"
	request.SSLPub = `-----BEGIN CERTIFICATE-----
	MIIFIjCCBAqgAwIBAgISBMqPy6+Q7TWmhdq4/ay5iYuGMA0GCSqGSIb3DQEBCwUA
	MDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD
	EwJSMzAeFw0yMTA0MjIwODM2MDZaFw0yMTA3MjEwODM2MDZaMBMxETAPBgNVBAMT
	CDhqMjRrLmNuMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy6oEUpeR
	RCx022ot/BLk/2xgZ2rEzdc1Jf2wmd7ggcr5iYoT/GNkRZhVXR5rwxvrqoFVXvEN
	9rucJDn/XXtRwtwWjXDnFVIW5Q1+o0lmAi2S/9vDJOKFnsN7eYIV9rxjerflhPYS
	z1hqovsKmevGCTJ8sf1xeYvjTvHjXkztfIzWE8+6pm7dkXEMlYC3o2NUdv1SASgv
	FT53/GHgs31S0laNpBrMbCTCUPj7o+F0BTNx5QJXODrf06R8duK3UmTR39JHOe7q
	7nceVbt3hrg7gsVg0D2PoXlP7GDIid36Drgmn+PtnAxJnFFL32I4e9F5PhxRTgms
	9cPQh0fiNz6fBwIDAQABo4ICTzCCAkswDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQW
	MBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBRc
	JZB6mjRUZk734Exw0/Uzqj8SvjAfBgNVHSMEGDAWgBQULrMXt1hWy65QCUDmH6+d
	ixTCxjBVBggrBgEFBQcBAQRJMEcwIQYIKwYBBQUHMAGGFWh0dHA6Ly9yMy5vLmxl
	bmNyLm9yZzAiBggrBgEFBQcwAoYWaHR0cDovL3IzLmkubGVuY3Iub3JnLzAfBgNV
	HREEGDAWggoqLjhqMjRrLmNuggg4ajI0ay5jbjBMBgNVHSAERTBDMAgGBmeBDAEC
	ATA3BgsrBgEEAYLfEwEBATAoMCYGCCsGAQUFBwIBFhpodHRwOi8vY3BzLmxldHNl
	bmNyeXB0Lm9yZzCCAQQGCisGAQQB1nkCBAIEgfUEgfIA8AB3AESUZS6w7s6vxEAH
	2Kj+KMDa5oK+2MsxtT/TM5a1toGoAAABePjvxWIAAAQDAEgwRgIhAJtE+BpvIeiN
	kpIqWeUn6GdsXkYvRZBjHhoUJyNfL3QMAiEAhixZSo8zOR9q0UhT78QH7eaW6iNC
	GjltNh5cmpF3Gi4AdQB9PvL4j/+IVWgkwsDKnlKJeSvFDngJfy5ql2iZfiLw1wAA
	AXj478V/AAAEAwBGMEQCIAlkCK2vUiF0EyvdlVWZojHyz8Tervt8lJcGbdR5P+qb
	AiBF6Xi5+B5XtEje7JRVSsgU2e2T+VaRumD+fVQ3qeyjQTANBgkqhkiG9w0BAQsF
	AAOCAQEASlIBffhriiokVvlb4pTDtGmqjvzqwN7UFVfr+FtHtjNsW3fAXgFJYqym
	lsBC0Y2pi8MK0tSpL6FQ+7hDf9VNh4IXa7gRSxKwmP6mXoNn0VAAKWjbvqN/MH8Z
	geGBnZ+NGVoGUvCaileZ+5kup5cUuPj/tzScGbUOczRxyU8BZ4KN0pJnHeCEGTqb
	+x8vWlFvMIi1DXWwlniQtEKH0+7HdFUTk/XYwezOyjRMJEFOWy3s3u6S/QZC2lYB
	G98ujn6T+dNG6D6QCiXeODQmO/uLm1NMQHm3EtHBHrVT4Z52EyYdRPGVbYjstCGO
	8j2FJNCJmepAV8rAZPek3cSl5u1Nvg==
	-----END CERTIFICATE-----
	-----BEGIN CERTIFICATE-----
	MIIEZTCCA02gAwIBAgIQQAF1BIMUpMghjISpDBbN3zANBgkqhkiG9w0BAQsFADA/
	MSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT
	DkRTVCBSb290IENBIFgzMB4XDTIwMTAwNzE5MjE0MFoXDTIxMDkyOTE5MjE0MFow
	MjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxCzAJBgNVBAMT
	AlIzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuwIVKMz2oJTTDxLs
	jVWSw/iC8ZmmekKIp10mqrUrucVMsa+Oa/l1yKPXD0eUFFU1V4yeqKI5GfWCPEKp
	Tm71O8Mu243AsFzzWTjn7c9p8FoLG77AlCQlh/o3cbMT5xys4Zvv2+Q7RVJFlqnB
	U840yFLuta7tj95gcOKlVKu2bQ6XpUA0ayvTvGbrZjR8+muLj1cpmfgwF126cm/7
	gcWt0oZYPRfH5wm78Sv3htzB2nFd1EbjzK0lwYi8YGd1ZrPxGPeiXOZT/zqItkel
	/xMY6pgJdz+dU/nPAeX1pnAXFK9jpP+Zs5Od3FOnBv5IhR2haa4ldbsTzFID9e1R
	oYvbFQIDAQABo4IBaDCCAWQwEgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8E
	BAMCAYYwSwYIKwYBBQUHAQEEPzA9MDsGCCsGAQUFBzAChi9odHRwOi8vYXBwcy5p
	ZGVudHJ1c3QuY29tL3Jvb3RzL2RzdHJvb3RjYXgzLnA3YzAfBgNVHSMEGDAWgBTE
	p7Gkeyxx+tvhS5B1/8QVYIWJEDBUBgNVHSAETTBLMAgGBmeBDAECATA/BgsrBgEE
	AYLfEwEBATAwMC4GCCsGAQUFBwIBFiJodHRwOi8vY3BzLnJvb3QteDEubGV0c2Vu
	Y3J5cHQub3JnMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly9jcmwuaWRlbnRydXN0
	LmNvbS9EU1RST09UQ0FYM0NSTC5jcmwwHQYDVR0OBBYEFBQusxe3WFbLrlAJQOYf
	r52LFMLGMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjANBgkqhkiG9w0B
	AQsFAAOCAQEA2UzgyfWEiDcx27sT4rP8i2tiEmxYt0l+PAK3qB8oYevO4C5z70kH
	ejWEHx2taPDY/laBL21/WKZuNTYQHHPD5b1tXgHXbnL7KqC401dk5VvCadTQsvd8
	S8MXjohyc9z9/G2948kLjmE6Flh9dDYrVYA9x2O+hEPGOaEOa1eePynBgPayvUfL
	qjBstzLhWVQLGAkXXmNs+5ZnPBxzDJOLxhF2JIbeQAcH5H0tZrUlo5ZYyOqA7s9p
	O5b85o3AM/OJ+CktFBQtfvBhcJVd9wvlwPsk+uyOy2HI7mNxKKgsBTt375teA2Tw
	UdHkhVNcsAKX1H7GNNLOEADksd86wuoXvg==
	-----END CERTIFICATE-----`
	request.SSLPri = `-----BEGIN PRIVATE KEY-----
	MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDLqgRSl5FELHTb
	ai38EuT/bGBnasTN1zUl/bCZ3uCByvmJihP8Y2RFmFVdHmvDG+uqgVVe8Q32u5wk
	Of9de1HC3BaNcOcVUhblDX6jSWYCLZL/28Mk4oWew3t5ghX2vGN6t+WE9hLPWGqi
	+wqZ68YJMnyx/XF5i+NO8eNeTO18jNYTz7qmbt2RcQyVgLejY1R2/VIBKC8VPnf8
	YeCzfVLSVo2kGsxsJMJQ+Puj4XQFM3HlAlc4Ot/TpHx24rdSZNHf0kc57urudx5V
	u3eGuDuCxWDQPY+heU/sYMiJ3foOuCaf4+2cDEmcUUvfYjh70Xk+HFFOCaz1w9CH
	R+I3Pp8HAgMBAAECggEBAJCnTTKsXI0jVkyRed/UO7n/a+mIOAhLloBJU6m2V5dY
	Zyx9WfylS47yx47AEKIfrp8IBW7Nn7FjR1+jofDvqf7Q+2OykS7fg0tk7lFjzIQ3
	gnUfzdbg3S8KgSSodJku7Sk4fFiXvnxuvC7OR5VTlPDrXw28iWBfLy/fNjLASN1w
	lTnOirLk0Aze0F4CGe8qYCxHb9whaHhF8Egd1/OQW4dz2m4gY+i82wixCPsMCELq
	2IFQD7G9ge6Jv568Bn0BNLR90ewfqr07FkhhFclr7sXZdTL8a/XAxMR2Q+rQPBTn
	POYxAqA8lI6NiPNxotsdKzD/T195XWl7IhTWd9+zKMECgYEA9p+9Q8kyezTGe5xH
	zkeT2w5mlXAuvJ9WQ8mGvln7tHpetefW0f8QgBdqUcgxzE4i9asfdyZhsirYv18d
	X+KI/Aoc1dmPQByX6aUFNy9aN1RkFwBs0XQc5ZXs0jTAZtS/5rQZF2R+ItVcTYmU
	JfD3+Hb2II9OLETBIgP1NvTkMQ8CgYEA02gsC89mvNsT3iQKVyXNO+hFoFJ9qBWa
	5abSbBqWi1wV1rZrm956jeoH/+q/8CQSxrw0OKYiuN4UJZWgwDYoVnPSfXFvn2XC
	gARgd1MRalsUj888CWxRSsjfwRbm76i8xut5TgP+lTznLlTHX0Uiouh8UExKZupK
	1xQ0N/7dwokCgYEAxIWPrLdjuZfZXeCvu2mAIXt9NTGanIWfmQ66h25uy2cDl5Jx
	tQ66GEgQWbDcNcBBzV+aCDkWeoH4C68AWkrK+4QXbDUcGGtwhicQW+Qo0JQ32bhl
	/hW3fR8WAn80nfyEGpOCtwhGFv3LuqGZ4w38mwo4oMNw+IelGrTqrus/9vMCgYEA
	pHMQcE9xt88jcVXPxWLed4+DOedxX+MOjz5pyTYVjWqggll8887ovqQifOz7Sa6F
	/JpCdR6nO+9k9KekWGDImBeY0YQANJpP5o2BDNFNZtJa+FNYfGwKAIapp+ZM1nkI
	ACUEQ6/Pxf+ORfgk7vM0skzvBveRmOUAOra4uCZGEyECgYBmZHxFvPu4hwvmxnDJ
	781Nlx/5vjO5Uj2eCzMhTx7Cia4ZLS6HaMkdEy3SAR2IXFhmkIPuZM9UwuD03JZO
	AFVLV0rX+GYdFmOJGSJY8rRjOVNzy/9FOSg9yqMoi6BcrHhtRSDdey+/GLCIdGw3
	8vZVCO2o4n3lftt5XnyoAEtwRQ==
	-----END PRIVATE KEY-----`

	response, err := client.BatchSetCdnDomainServerCertificate(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return response.String()
}
