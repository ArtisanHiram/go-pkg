package __obf_2bca6d09a5e8d9c4

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_a25b1325f319cddb func(string, string, []string) (string, string, error)

var __obf_b1de0b5898cda430 = map[string]__obf_a25b1325f319cddb{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_2bca6d09a5e8d9c4 string) bool {
	_, __obf_674eff4418cbc62c := __obf_b1de0b5898cda430[__obf_2bca6d09a5e8d9c4]
	return __obf_674eff4418cbc62c
}

func Run(__obf_2bca6d09a5e8d9c4 string, __obf_1ea2f422bf0146d8, __obf_943ae9bbc13cc8d0 string, __obf_28454e8a224b7ff8 []string) (string, string, error) {
	return __obf_b1de0b5898cda430[__obf_2bca6d09a5e8d9c4](__obf_1ea2f422bf0146d8, __obf_943ae9bbc13cc8d0, __obf_28454e8a224b7ff8)
}
