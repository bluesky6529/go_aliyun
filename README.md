# go_aliyun

## before
> go mod init github.com/bluesky6529/go_aliyun

## Must be do 
+ use govendor
> go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
> 
> go get -u github.com/spf13/cobra
>
> go gt -u github.com/spf13/viper

#### This project use cobra & viper & aliyun go sdk to complete.

### 完成的功能
```
Usage:
   [command]

Available Commands:
  BatchSetCdnDomainServerCertificate 更新CDN 憑證
  BatchSetDcdnDomainCertificate      更新DCDN 憑證
  DescribeDcdnDomainCertificateInfo  查看DCDN 域名憑證
  DescribeDcdnRefreshTasks           刷新DCDN 進度
  DescribeDomainCertificateInfo      查看CDN 域名憑證
  DescribeRefreshTasks               刷新CDN 進度
  QueryAccountBalance                查詢帳戶餘額
  RefreshDcdnObjectCaches            刷新DCDN url/目錄
  RefreshObjectCaches                刷新CDN url/目錄
  help                               Help about any command

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

### for myself memo
#### remember windows to linux env
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build .
```