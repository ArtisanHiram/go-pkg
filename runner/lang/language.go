package __obf_870ef91d695d9345

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_bb47752e003a294b func(string, string, []string) (string, string, error)

var __obf_fc849539a5e051ef = map[string]__obf_bb47752e003a294b{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_870ef91d695d9345 string) bool {
	_, __obf_9f50fa52264d2644 := __obf_fc849539a5e051ef[__obf_870ef91d695d9345]
	return __obf_9f50fa52264d2644
}

func Run(__obf_870ef91d695d9345 string, __obf_8c3fd997f4d44f4f, __obf_1ef3dd53868de0d8 string, __obf_35314af6dd6e1909 []string) (string, string, error) {
	return __obf_fc849539a5e051ef[__obf_870ef91d695d9345](__obf_8c3fd997f4d44f4f, __obf_1ef3dd53868de0d8, __obf_35314af6dd6e1909)
}
