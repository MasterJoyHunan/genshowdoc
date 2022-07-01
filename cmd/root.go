package cmd

import (
	"fmt"
	"os"

	"github.com/MasterJoyHunan/genshowdoc/generator"
	"github.com/MasterJoyHunan/genshowdoc/prepare"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
)

var (
	rootCmd = &cobra.Command{
		Use:   "genrpc",
		Short: "生成 GRPC 的项目结构",
		Args:  cobra.ExactValidArgs(1),
		RunE:  GenShowDoc,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&prepare.Api, "api", "", "showdoc 的 api 接口")
	rootCmd.Flags().StringVar(&prepare.Key, "key", "", "showdoc 项目的 key")
	rootCmd.Flags().StringVar(&prepare.Token, "token", "", "showdoc 项目的 token")
}

func GenShowDoc(cmd *cobra.Command, args []string) error {
	prepare.ApiFile = args[0]
	apiInfo, err := parser.Parse(args[0])
	if err != nil {
		return err
	}
	prepare.ApiInfo = apiInfo

	return generator.GenShowDoc()
}
