package __obf_37db86f540013a3c

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_7e30183e575a1afa func(string, string, []string) (string, string, error)

var __obf_904ba9c666b5c584 = map[string]__obf_7e30183e575a1afa{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_37db86f540013a3c string) bool {
	_, __obf_3c88eefeb3265f13 := __obf_904ba9c666b5c584[__obf_37db86f540013a3c]
	return __obf_3c88eefeb3265f13
}

func Run(__obf_37db86f540013a3c string, __obf_8de47fcd4a7145d3, __obf_009ffdaa71ae5da1 string, __obf_f7e7ba44300171d6 []string) (string, string, error) {
	return __obf_904ba9c666b5c584[__obf_37db86f540013a3c](__obf_8de47fcd4a7145d3, __obf_009ffdaa71ae5da1, __obf_f7e7ba44300171d6)
}
