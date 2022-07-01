package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/MasterJoyHunan/genshowdoc/prepare"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func GenResponseExample(tp spec.Type) string {
	var sb strings.Builder
	sb.WriteString("{")
	deepGenResponseExample(&sb, tp)
	sb.WriteString("}")

	// 格式化
	var bb bytes.Buffer
	err := json.Indent(&bb, []byte(sb.String()), "", "    ")
	if err != nil {
		panic(err)
	}
	return bb.String()
}

func deepGenResponseExample(sb *strings.Builder, tp spec.Type) {
	defineStruct, ok := tp.(spec.DefineStruct)
	if !ok {
		return
	}
	for _, t := range prepare.ApiInfo.Types {
		if t.Name() == tp.Name() {
			defineStruct = t.(spec.DefineStruct)
		}
	}
	for i, m := range defineStruct.Members {
		comma := ","
		if i == len(defineStruct.Members)-1 {
			comma = ""
		}
		switch v := m.Type.(type) {
		case spec.PrimitiveType:
			sb.WriteString(fmt.Sprintf("\"%s\" : \"%s\"%s", getMemberName(m.Name, m.Tag), v.RawName, comma))
		case spec.MapType:
			sb.WriteString(fmt.Sprintf("\"%s\" : {", getMemberName(m.Name, m.Tag)))
			if mapInternalType, isPrimitiveType := v.Value.(spec.PrimitiveType); isPrimitiveType {
				sb.WriteString(fmt.Sprintf("\"map_key_is_%s\" : \"map_value_is_%s\",", v.Key, mapInternalType.RawName))
				sb.WriteString(fmt.Sprintf("\"map_key_is_%s\" : \"map_value_is_%s\"", v.Key, mapInternalType.RawName))
			} else {
				sb.WriteString(fmt.Sprintf("\"map_key_is_%s\" : {", v.Key))
				deepGenResponseExample(sb, v.Value)
				sb.WriteString("}")
			}
			sb.WriteString("}" + comma)
		case spec.ArrayType:
			// 可能是多维数组
			typeName := m.Type.Name()
			for {
				typeName := strings.TrimPrefix(typeName, "[]")
				sb.WriteString(fmt.Sprintf("\"%s\" : [", getMemberName(m.Name, m.Tag)))
				if !strings.HasPrefix(typeName, "[]") {
					arrayInternalType, isPrimitiveType := v.Value.(spec.PrimitiveType)
					if isPrimitiveType {
						sb.WriteString(fmt.Sprintf("\"%s\"", arrayInternalType.Name()))
						sb.WriteString(",")
						sb.WriteString(fmt.Sprintf("\"%s\"", arrayInternalType.Name()))
					} else {
						sb.WriteString("{")
						deepGenResponseExample(sb, v.Value)
						sb.WriteString("}")
					}
					sb.WriteString("]" + comma)
					break
				}
				sb.WriteString("]")
			}
		case spec.DefineStruct:
			sb.WriteString(fmt.Sprintf("\"%s\": {", getMemberName(m.Name, m.Tag)))
			deepGenResponseExample(sb, m.Type)
			sb.WriteString("}" + comma)
		}
	}
}
