package cmd

import (
	"fmt"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/DescribeDcdnRefreshTasks"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var describedcdnrefreshtasks = &cobra.Command{
	Use:   "DescribeDcdnRefreshTasks",
	Short: "刷新DCDN 進度",
	Long:  "用來查看刷新DCDN進度，請輸入你刷新的url/目錄，使用方法如下 : \nDescribeDcdnRefreshTasks -a <帳戶名稱> -u <url> \nex: DescribeDcdnDomainCertificateInfo -a account -u https://abc.com/",
	Run: func(cmd *cobra.Command, args []string) {
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request := DescribeDcdnRefreshTasks.DescribeDcdnRefreshTasks(AccessKeyID, AccessKeySecret, url)
		fmt.Printf("%s", request)
	},
}

func init() {
	describedcdnrefreshtasks.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	describedcdnrefreshtasks.MarkFlagRequired("account")
	describedcdnrefreshtasks.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:https://abc.com/)")
	//describedcdnrefreshtasks.MarkFlagRequired("url")
}
