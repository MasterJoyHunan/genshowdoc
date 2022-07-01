package generator

import (
	"fmt"
	"strings"

	"github.com/MasterJoyHunan/genshowdoc/prepare"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const memo = "更多返回错误代码请看首页的错误代码描述"

func GenShowDoc() error {
	for _, g := range prepare.ApiInfo.Service.Groups {
		for _, r := range g.Routes {
			api, err := MakeApiDoc(r, g)
			if err != nil {
				return err
			}
			CallApi(api)
		}
	}
	return nil
}

func MakeApiDoc(route spec.Route, group spec.Group) (api ApiDoc, err error) {
	var req []Request
	var resp []Response
	makeRequest(route.RequestType, &req, "")
	makeResponse(route.ResponseType, &resp, "")
	api.Desc = makeTitle(route)
	api.Url = group.GetAnnotation("prefix") + route.Path
	api.Method = route.Method
	api.Request = req
	api.Response = resp
	api.ResponseExample = GenResponseExample(route.ResponseType)
	api.Memo = memo
	if group.GetAnnotation("group") != "" {
		api.Group = group.GetAnnotation("group")
	}
	if group.GetAnnotation("swtags") != "" {
		api.Group = group.GetAnnotation("swtags")
	}
	return
}

func makeTitle(route spec.Route) string {
	if route.AtDoc.Text != "" {
		return strings.Trim(route.AtDoc.Text, "\"")
	}
	return route.Handler
}

func makeRequest(tp spec.Type, req *[]Request, prefix string) {
	if tp == nil {
		return
	}

	defineStruct, ok := tp.(spec.DefineStruct)
	if !ok {
		return
	}

	for _, t := range prepare.ApiInfo.Types {
		if t.Name() == tp.Name() {
			defineStruct = t.(spec.DefineStruct)
		}
	}

	for _, m := range defineStruct.Members {
		current := Request{
			Name:      prefix + getMemberName(m.Name, m.Tag),
			IsRequire: getIsRequire(m.Tag),
			Type:      getMemberType(m.Type),
			Memo:      strings.TrimPrefix(m.GetComment(), "//"),
		}
		*req = append(*req, current)

		switch v := m.Type.(type) {
		case spec.MapType:
			makeRequest(v.Value, req, current.Name+".")
		case spec.ArrayType:
			makeRequest(v.Value, req, current.Name+".")
		case spec.DefineStruct:
			makeRequest(m.Type, req, current.Name+".")
		}
	}
}

func makeResponse(tp spec.Type, resp *[]Response, prefix string) {
	if tp == nil {
		return
	}

	defineStruct, ok := tp.(spec.DefineStruct)
	if !ok {
		return
	}
	for _, t := range prepare.ApiInfo.Types {
		if t.Name() == tp.Name() {
			defineStruct = t.(spec.DefineStruct)
		}
	}
	for _, m := range defineStruct.Members {
		current := Response{
			Name: prefix + getMemberName(m.Name, m.Tag),
			Type: getMemberType(m.Type),
			Memo: strings.TrimPrefix(m.GetComment(), "//"),
		}
		*resp = append(*resp, current)

		switch v := m.Type.(type) {
		case spec.MapType:
			makeResponse(v.Value, resp, current.Name+".")
		case spec.ArrayType:
			makeResponse(v.Value, resp, current.Name+".")
		case spec.DefineStruct:
			makeResponse(m.Type, resp, current.Name+".")
		}
	}
}

func getMemberName(defaultTypeName, tag string) string {
	m := parseTag(tag)
	if v, ok := m["path"]; ok {
		return v
	}
	if v, ok := m["header"]; ok {
		return v
	}
	if v, ok := m["form"]; ok {
		return v
	}
	if v, ok := m["json"]; ok {
		return v
	}
	fmt.Println(tag + " 没有匹配的参数名")
	return defaultTypeName
}

func getIsRequire(tag string) string {
	m := parseTag(tag)
	for _, v := range m {
		if strings.Contains(v, "optional") {
			return "否"
		}
		if strings.Contains(v, "required") {
			return "是"
		}
	}
	return "是"
}

func getMemberType(t spec.Type) string {
	switch v := t.(type) {
	case spec.PrimitiveType:
		return v.Name()
	case spec.MapType:
		return "map"
	case spec.ArrayType:
		return "array"
	case spec.DefineStruct:
		return "object"
	}
	fmt.Println(t.Name() + " 没有匹配的参数类型")
	return "UNKNOW"
}

func parseTag(tag string) map[string]string {
	tagMap := make(map[string]string)
	tag = strings.Trim(tag, "`")
	tags := strings.Split(tag, " ")
	for _, t := range tags {
		tagNameAndValue := strings.Split(t, ":")
		tagMap[tagNameAndValue[0]] = strings.Trim(tagNameAndValue[1], "\"")
	}
	return tagMap
}
