package __obf_10244cbc933db822

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_a7f45821c25810a1 func(string, string, []string) (string, string, error)

var __obf_f208e66b06270200 = map[string]__obf_a7f45821c25810a1{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_10244cbc933db822 string) bool {
	_, __obf_b8b31bbbaa61db59 := __obf_f208e66b06270200[__obf_10244cbc933db822]
	return __obf_b8b31bbbaa61db59
}

func Run(__obf_10244cbc933db822 string, __obf_604a6a6619d19fa4, __obf_9f3d0b3c4936753f string, __obf_154447dc8d89cf79 []string) (string, string, error) {
	return __obf_f208e66b06270200[__obf_10244cbc933db822](__obf_604a6a6619d19fa4, __obf_9f3d0b3c4936753f, __obf_154447dc8d89cf79)
}
