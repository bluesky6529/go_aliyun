package cmd

import (
	"fmt"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/DescribeDomainCertificateInfo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var describedomaincertificateinfo = &cobra.Command{
	Use:   "DescribeDomainCertificateInfo",
	Short: "查看域名憑證",
	Long:  "查看域名憑證，使用方法如下 : \nDescribeDomainCertificateInfo -a <帳戶名稱> -u <url> \nex: DescribeRefreshTasks -a account -u abc.com",
	Run: func(cmd *cobra.Command, args []string) {
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request := DescribeDomainCertificateInfo.DescribeDomainCertificateInfo(AccessKeyID, AccessKeySecret, url)
		fmt.Printf("%s", request)
	},
}

func init() {
	describerefreshtasks.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	describerefreshtasks.MarkFlagRequired("account")
	describerefreshtasks.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:abc.com)")
	//describerefreshtasks.MarkFlagRequired("url")
}
