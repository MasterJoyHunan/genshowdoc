package test

import (
	"testing"

	"github.com/MasterJoyHunan/genshowdoc/generator"
	"github.com/MasterJoyHunan/genshowdoc/prepare"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
)

func TestGenShowDoc(t *testing.T) {
	prepare.ApiFile = "api/someapp.api"
	apiInfo, err := parser.Parse(prepare.ApiFile)
	if err != nil {
		t.Failed()
	}
	prepare.ApiInfo = apiInfo
	prepare.Api = "http://127.0.0.1:8987/server/index.php?s=/api/item/updateByApi"
	prepare.Key = "094bbed23398a4a05886404eae8229d2630441252"
	prepare.Token = "a2bc803851c08b02afb7414193aa652b1413697826"
	if err := generator.GenShowDoc(); err != nil {
		t.Failed()
	}
}
