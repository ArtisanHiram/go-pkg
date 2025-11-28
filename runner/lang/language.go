package __obf_e1a916ea6db25f42

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_c373162f5d3705c2 func(string, string, []string) (string, string, error)

var __obf_025c974994ef6ef3 = map[string]__obf_c373162f5d3705c2{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_e1a916ea6db25f42 string) bool {
	_, __obf_87d8e8c11b7f05be := __obf_025c974994ef6ef3[__obf_e1a916ea6db25f42]
	return __obf_87d8e8c11b7f05be
}

func Run(__obf_e1a916ea6db25f42 string, __obf_c154cebfa84edf49, __obf_c9403dbe407ddb7f string, __obf_0a696dd30c676e19 []string) (string, string, error) {
	return __obf_025c974994ef6ef3[__obf_e1a916ea6db25f42](__obf_c154cebfa84edf49, __obf_c9403dbe407ddb7f, __obf_0a696dd30c676e19)
}
