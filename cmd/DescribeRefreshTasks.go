package cmd

import (
	"fmt"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/DescribeRefreshTasks"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var describerefreshtasks = &cobra.Command{
	Use:   "DescribeRefreshTasks",
	Short: "刷新CDN 進度",
	Long:  "用來查看刷新CDN 進度，使用方法如下 : \nDescribeRefreshTasks -a <帳戶名稱> -u <url> \nex: DescribeRefreshTasks -a account -u https://abc.com",
	Run: func(cmd *cobra.Command, args []string) {
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request := DescribeRefreshTasks.DescribeRefreshTasks(AccessKeyID, AccessKeySecret, url)
		fmt.Printf("%s", request)
	},
}

func init() {
	describerefreshtasks.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	describerefreshtasks.MarkFlagRequired("account")
	describerefreshtasks.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:https://abc.com/)")
	//describerefreshtasks.MarkFlagRequired("url")
}
