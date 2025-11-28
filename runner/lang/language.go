package __obf_a42068b4935a1fce

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_437b415f53385194 func(string, string, []string) (string, string, error)

var __obf_1339a07926ae0805 = map[string]__obf_437b415f53385194{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_a42068b4935a1fce string) bool {
	_, __obf_3692e413a8897ddf := __obf_1339a07926ae0805[__obf_a42068b4935a1fce]
	return __obf_3692e413a8897ddf
}

func Run(__obf_a42068b4935a1fce string, __obf_d799128e96aad1f2, __obf_89d9794bc2b85dab string, __obf_bdc3fb6d471f4696 []string) (string, string, error) {
	return __obf_1339a07926ae0805[__obf_a42068b4935a1fce](__obf_d799128e96aad1f2, __obf_89d9794bc2b85dab, __obf_bdc3fb6d471f4696)
}
