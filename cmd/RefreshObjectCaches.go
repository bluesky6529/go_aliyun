package cmd

import (
	"log"
	"strings"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/RefreshObjectCaches"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var refreshobjectcaches = &cobra.Command{
	Use:   "RefreshObjectCaches",
	Short: "刷新CDN url/目錄",
	Long:  "用來刷新cdn目錄 or url 使用方法如下 : \n刷新目錄: RefreshObjectCaches -a <account> -d Directory -u http://abc.com/ \n刷新url: RefreshObjectCaches -a <account> -d File -u http://abc.com/index.html",
	Run: func(cmd *cobra.Command, args []string) {
		var request string
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request = RefreshObjectCaches.RefreshObjectCaches(AccessKeyID, AccessKeySecret, strings.Split(strings.Trim(url, "'"), ","), dir)
		log.Printf("输出结果: %s", request)
	},
}

func init() {
	refreshobjectcaches.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	refreshobjectcaches.Flags().StringVarP(&url, "url", "u", "", "輸入刷新url(EX:https://abc.com/) (require)")
	refreshobjectcaches.Flags().StringVarP(&dir, "dir", "d", "", "刷新目錄輸入Directory 刷新url 輸入File (require)")
	refreshobjectcaches.MarkFlagRequired("url")
	refreshobjectcaches.MarkFlagRequired("account")
	refreshobjectcaches.MarkFlagRequired("dir")
}
