package cmd

import (
	"log"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/BatchSetCdnDomainServerCertificate"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var batchsetcdndomainservercertificate = &cobra.Command{
	Use:   "BatchSetCdnDomainServerCertificate",
	Short: "更新憑證",
	Long:  "用來更新憑證 使用方法如下 : \n刷新目錄: BatchSetCdnDomainServerCertificate -a <account> \n刷新url: BatchSetCdnDomainServerCertificate -a <account> ",
	Run: func(cmd *cobra.Command, args []string) {
		var request string
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request = BatchSetCdnDomainServerCertificate.BatchSetCdnDomainServerCertificate(AccessKeyID, AccessKeySecret)
		log.Printf("输出结果: %s", request)
	},
}

func init() {
	batchsetcdndomainservercertificate.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	//batchsetcdndomainservercertificate.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:https://abc.com/) (require)")
	//batchsetcdndomainservercertificate.Flags().StringVarP(&dir, "dir", "d", "", "刷新目錄輸入Directory 刷新url 輸入File (require)")
	//batchsetcdndomainservercertificate.MarkFlagRequired("url")
	batchsetcdndomainservercertificate.MarkFlagRequired("account")
	//batchsetcdndomainservercertificate.MarkFlagRequired("dir")
}
