package generator

import (
	"fmt"
	"strings"
)

type ApiDoc struct {
	sb              strings.Builder
	Desc            string     // 简要描述,标题
	Url             string     // 请求URL
	Method          string     // 请求方式
	Request         []Request  // 请求参数
	Response        []Response // 返回示例
	ResponseExample string     // 返回示例
	Memo            string     // 备注
	Group           string     // 所属
}

type Request struct {
	Name      string // 参数名
	IsRequire string // 是否必填
	Type      string // 类型
	Memo      string // 说明
}

type Response struct {
	Name string // 参数名
	Type string // 类型
	Memo string // 说明
}

func (d *ApiDoc) makeDesc() {
	d.sb.WriteString("##### 简要描述\n\n")
	d.sb.WriteString("- " + d.Desc)
	d.sb.WriteString("\n\n")
	d.sb.WriteString("- 该文档由工具自动生成，请勿修改，重复生成将覆盖该文档\n\n")
}

func (d *ApiDoc) makeUrl() {
	d.sb.WriteString("##### 请求URL\n\n")
	d.sb.WriteString("`" + d.Url + "`\n\n")
}

func (d *ApiDoc) makeMethod() {
	d.sb.WriteString("##### 请求方式 \n\n")
	d.sb.WriteString("- " + strings.ToUpper(d.Method))
	d.sb.WriteString("\n\n")
}

func (d *ApiDoc) makeRequest() {
	if len(d.Request) == 0 {
		return
	}
	d.sb.WriteString("##### 参数\n\n")
	d.sb.WriteString("|参数名|必选|类型|说明|\n")
	d.sb.WriteString("|:-|:-|:-|:-|\n")

	for _, r := range d.Request {
		reqMemo := r.Memo
		if reqMemo == "" {
			reqMemo = "暂无描述"
		}
		d.sb.WriteString(fmt.Sprintf("|%s|%s|%s|%s|\n", r.Name, r.IsRequire, r.Type, reqMemo))
	}
	d.sb.WriteString("\n")
}

func (d *ApiDoc) makeResponse() {
	if len(d.Response) == 0 {
		return
	}
	d.sb.WriteString("##### 返回参数说明 \n\n")
	d.sb.WriteString("|参数名|类型|说明|\n")
	d.sb.WriteString("|:-|:-|:-|\n")
	for _, r := range d.Response {
		respMemo := r.Memo
		if respMemo == "" {
			respMemo = "暂无描述"
		}
		d.sb.WriteString(fmt.Sprintf("|%s|%s|%s|\n", r.Name, r.Type, respMemo))
	}
	d.sb.WriteString("\n")
}

func (d *ApiDoc) makeResponseExample() {
	if len(d.Response) == 0 {
		return
	}
	d.sb.WriteString("##### 返回示例 \n\n")
	d.sb.WriteString("```\n")
	d.sb.WriteString(d.ResponseExample)
	d.sb.WriteString("```\n")
}

func (d *ApiDoc) makeMemo() {
	d.sb.WriteString("##### 备注 \n\n")
	d.sb.WriteString("- ")
	d.sb.WriteString(d.Memo)
	d.sb.WriteString("\n")
}

func (d *ApiDoc) MakeDoc() string {
	d.makeDesc()
	d.makeUrl()
	d.makeMethod()
	d.makeRequest()
	d.makeResponse()
	d.makeResponseExample()
	d.makeMemo()
	return d.sb.String()
}
