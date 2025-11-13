package __obf_ba7d5cb9876e322e

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_8613fa2957c00822 func(string, string, []string) (string, string, error)

var __obf_4cc7243fb107deb6 = map[string]__obf_8613fa2957c00822{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_ba7d5cb9876e322e string) bool {
	_, __obf_4203d354401f885c := __obf_4cc7243fb107deb6[__obf_ba7d5cb9876e322e]
	return __obf_4203d354401f885c
}

func Run(__obf_ba7d5cb9876e322e string, __obf_518c2758de91cae3, __obf_0a136912cddcee1a string, __obf_e21fcb328c394b3c []string) (string, string, error) {
	return __obf_4cc7243fb107deb6[__obf_ba7d5cb9876e322e](__obf_518c2758de91cae3, __obf_0a136912cddcee1a, __obf_e21fcb328c394b3c)
}
