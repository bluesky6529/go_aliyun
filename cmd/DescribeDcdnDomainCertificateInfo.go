package cmd

import (
	"fmt"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/DescribeDcdnDomainCertificateInfo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var describedcdndomaincertificateInfo = &cobra.Command{
	Use:   "DescribeDcdnDomainCertificateInfo",
	Short: "查看DCDN 域名憑證",
	Long:  "查看DCDN 域名憑證，使用方法如下 : \nDescribeDcdnDomainCertificateInfo -a <帳戶名稱> -u <url> \nex: DescribeDcdnDomainCertificateInfo -a account -u abc.com",
	Run: func(cmd *cobra.Command, args []string) {
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request := DescribeDcdnDomainCertificateInfo.DescribeDcdnDomainCertificateInfo(AccessKeyID, AccessKeySecret, url)
		fmt.Printf("%s", request)
	},
}

func init() {
	describedcdndomaincertificateInfo.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	describedcdndomaincertificateInfo.MarkFlagRequired("account")
	describedcdndomaincertificateInfo.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:abc.com)")
	describedcdndomaincertificateInfo.MarkFlagRequired("url")
}
