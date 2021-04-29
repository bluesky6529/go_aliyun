package cmd

import (
	"fmt"
	"io/ioutil"
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

		sslpub_val, err := ioutil.ReadFile(sslpub)
		if err != nil {
			fmt.Println("read fail", err)
		}
		//var sslpub_str = string(sslpub_val)

		sslpri_val, err := ioutil.ReadFile(sslpri)
		if err != nil {
			fmt.Println("read fail", err)
		}
		//var sslpri_str = string(sslpri_val)

		request = BatchSetCdnDomainServerCertificate.BatchSetCdnDomainServerCertificate(AccessKeyID, AccessKeySecret, url, string(sslpub_val), string(sslpri_val))
		log.Printf("输出结果: %s", request)
	},
}

func init() {
	batchsetcdndomainservercertificate.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	batchsetcdndomainservercertificate.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:https://abc.com/) (require)")
	batchsetcdndomainservercertificate.Flags().StringVarP(&sslpub, "sslpub", "k", "", "(require)")
	batchsetcdndomainservercertificate.Flags().StringVarP(&sslpri, "sslpri", "j", "", "(require)")
	batchsetcdndomainservercertificate.MarkFlagRequired("url")
	batchsetcdndomainservercertificate.MarkFlagRequired("account")
	batchsetcdndomainservercertificate.MarkFlagRequired("sslpub")
	batchsetcdndomainservercertificate.MarkFlagRequired("sslpri")
}
