package __obf_0a6a4c2205ce53b4

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_1d24fcbb13a36e32 func(string, string, []string) (string, string, error)

var __obf_200581d6d1f237f6 = map[string]__obf_1d24fcbb13a36e32{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_0a6a4c2205ce53b4 string) bool {
	_, __obf_464ddda7bb91b1a3 := __obf_200581d6d1f237f6[__obf_0a6a4c2205ce53b4]
	return __obf_464ddda7bb91b1a3
}

func Run(__obf_0a6a4c2205ce53b4 string, __obf_2ff54b0509d918d5, __obf_71b9dfb4b13a29cc string, __obf_a342349a71a1d077 []string) (string, string, error) {
	return __obf_200581d6d1f237f6[__obf_0a6a4c2205ce53b4](__obf_2ff54b0509d918d5, __obf_71b9dfb4b13a29cc, __obf_a342349a71a1d077)
}
