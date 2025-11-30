package __obf_9f8d52b027865418

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_83e1f09c573b71ba func(string, string, []string) (string, string, error)

var __obf_69093de07c877f4c = map[string]__obf_83e1f09c573b71ba{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_9f8d52b027865418 string) bool {
	_, __obf_8efaf652a594c6f7 := __obf_69093de07c877f4c[__obf_9f8d52b027865418]
	return __obf_8efaf652a594c6f7
}

func Run(__obf_9f8d52b027865418 string, __obf_8c79a76513d07503, __obf_2dd9bbd829c254e3 string, __obf_1d3ec9b2b75f6c72 []string) (string, string, error) {
	return __obf_69093de07c877f4c[__obf_9f8d52b027865418](__obf_8c79a76513d07503, __obf_2dd9bbd829c254e3, __obf_1d3ec9b2b75f6c72)
}
