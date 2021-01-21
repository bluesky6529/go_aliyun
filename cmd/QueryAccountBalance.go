package cmd

import (
	"log"

	_ "github.com/bluesky6529/go_aliyun/configs"
	"github.com/bluesky6529/go_aliyun/internal/QueryAccountBalance"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var queryaccountbalance = &cobra.Command{
	Use:   "QueryAccountBalance",
	Short: "查詢帳戶餘額",
	Long:  "用來查詢帳戶餘額使用方法如下 : \nQueryAccountBalance -a <帳戶名稱> ex: QueryAccountBalance -a account",
	Run: func(cmd *cobra.Command, args []string) {
		var AccessKeyID = viper.GetString(account + ".AccessKeyID")
		var AccessKeySecret = viper.GetString(account + ".AccessKeySecret")

		request := QueryAccountBalance.QueryAccountBalance(AccessKeyID, AccessKeySecret)
		log.Printf("%s 帳戶餘額: %s", account, request)
	},
}

func init() {
	queryaccountbalance.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	queryaccountbalance.MarkFlagRequired("account")
}
