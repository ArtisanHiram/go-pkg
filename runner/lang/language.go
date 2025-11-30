package __obf_097dcda73a8f146b

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_1a3145c84fb7aedb func(string, string, []string) (string, string, error)

var __obf_18052c42b1f72015 = map[string]__obf_1a3145c84fb7aedb{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_097dcda73a8f146b string) bool {
	_, __obf_f9e99382294ee51c := __obf_18052c42b1f72015[__obf_097dcda73a8f146b]
	return __obf_f9e99382294ee51c
}

func Run(__obf_097dcda73a8f146b string, __obf_f1555f62b28d62a5, __obf_773cbd1e367c3b6d string, __obf_432a95a38ed6a5b1 []string) (string, string, error) {
	return __obf_18052c42b1f72015[__obf_097dcda73a8f146b](__obf_f1555f62b28d62a5, __obf_773cbd1e367c3b6d, __obf_432a95a38ed6a5b1)
}
