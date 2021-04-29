package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/BatchSetDcdnDomainCertificate"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var batchsetdcdndomaincertificate = &cobra.Command{
	Use:   "BatchSetDcdnDomainCertificate",
	Short: "更新DCDN 憑證",
	Long:  "用來更新憑證 使用方法如下 : \n BatchSetDcdnDomainCertificate -a <account> -u <url> -k 公鑰位置 -j 私鑰位置\n EX: BatchSetDcdnDomainCertificate -a <account> -u abc.com -k fillchain.pem -j prikey.pem",
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

		request = BatchSetDcdnDomainCertificate.BatchSetDcdnDomainCertificate(AccessKeyID, AccessKeySecret, url, string(sslpub_val), string(sslpri_val))
		log.Printf("输出结果: %s", request)
	},
}

func init() {
	batchsetdcdndomaincertificate.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	batchsetdcdndomaincertificate.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:https://abc.com/) (require)")
	batchsetdcdndomaincertificate.Flags().StringVarP(&sslpub, "sslpub", "k", "", "(require)")
	batchsetdcdndomaincertificate.Flags().StringVarP(&sslpri, "sslpri", "j", "", "(require)")
	batchsetdcdndomaincertificate.MarkFlagRequired("url")
	batchsetdcdndomaincertificate.MarkFlagRequired("account")
	batchsetdcdndomaincertificate.MarkFlagRequired("sslpub")
	batchsetdcdndomaincertificate.MarkFlagRequired("sslpri")
}
