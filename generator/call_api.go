package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MasterJoyHunan/genshowdoc/prepare"
)

type showDocApiParam struct {
	ApiKey      string `json:"api_key"`
	ApiToken    string `json:"api_token"`
	CatName     string `json:"cat_name"`
	PageTitle   string `json:"page_title"`
	PageContent string `json:"page_content"`
}

type callApiResponse struct {
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

func CallApi(api ApiDoc) {
	req := showDocApiParam{
		ApiKey:      prepare.Key,
		ApiToken:    prepare.Token,
		CatName:     getCatName(api.Group),
		PageTitle:   api.Desc,
		PageContent: api.MakeDoc(),
	}

	marshal, err := json.Marshal(&req)
	if err != nil {
		fmt.Println("生成返回结构体示例错误")
		os.Exit(1)
	}
	resp, err := http.Post(prepare.Api, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		fmt.Println("发送接口错误" + err.Error())
		os.Exit(1)
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请求api错误" + err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	var apiResponse callApiResponse

	err = json.Unmarshal(all, &apiResponse)
	if err != nil {
		fmt.Println("解析返回错误" + err.Error())
		os.Exit(1)
	}
	if apiResponse.ErrorCode != 0 {
		fmt.Println(apiResponse.ErrorMessage)
		os.Exit(1)
	}
	fmt.Println("成功生成【 " + api.Desc + " 】接口文档")
}

func getCatName(catName string) string {
	catName = strings.Trim(catName, "\"")
	cat := ""
	title, ok := prepare.ApiInfo.Info.Properties["title"]
	if ok {
		cat = strings.Trim(title, "\"") + "/" + catName
	}
	return cat
}
