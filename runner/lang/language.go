package __obf_f5414d66650d478d

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_6c6a91b99e4ee997 func(string, string, []string) (string, string, error)

var __obf_7f8a440cb6497fe4 = map[string]__obf_6c6a91b99e4ee997{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_f5414d66650d478d string) bool {
	_, __obf_f458bf312d2856ad := __obf_7f8a440cb6497fe4[__obf_f5414d66650d478d]
	return __obf_f458bf312d2856ad
}

func Run(__obf_f5414d66650d478d string, __obf_0fcb0fbbfcd6733f, __obf_c6d66cce6a41e12a string, __obf_b114cd8343da370e []string) (string, string, error) {
	return __obf_7f8a440cb6497fe4[__obf_f5414d66650d478d](__obf_0fcb0fbbfcd6733f, __obf_c6d66cce6a41e12a, __obf_b114cd8343da370e)
}
