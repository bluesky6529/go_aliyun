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
  DescribeDcdnRefreshTasks 刷新cdn 進度
  DescribeRefreshTasks     刷新dcdn 進度
  QueryAccountBalance      查詢帳戶餘額
  RefreshDcdnObjectCaches  刷新Dcdn url/目錄
  RefreshObjectCaches      刷新CDN url/目錄
  help                     Help about any command

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```
