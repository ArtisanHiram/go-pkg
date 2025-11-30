package __obf_ce68e709ffad86c7

import (
	cpp "github.com/ArtisanHiram/go-pkg/runner/lang/cpp"
	golang "github.com/ArtisanHiram/go-pkg/runner/lang/golang"
	groovy "github.com/ArtisanHiram/go-pkg/runner/lang/groovy"
	java "github.com/ArtisanHiram/go-pkg/runner/lang/java"
	javascript "github.com/ArtisanHiram/go-pkg/runner/lang/javascript"
	python "github.com/ArtisanHiram/go-pkg/runner/lang/python"
	typescript "github.com/ArtisanHiram/go-pkg/runner/lang/typescript"
)

type __obf_19952cd758271b88 func(string, string, []string) (string, string, error)

var __obf_32414073ae402561 = map[string]__obf_19952cd758271b88{
	"cpp":        cpp.Run,
	"go":         golang.Run,
	"groovy":     groovy.Run,
	"java":       java.Run,
	"javascript": javascript.Run,
	"python":     python.Run,
	"typescript": typescript.Run,
}

func IsSupported(__obf_ce68e709ffad86c7 string) bool {
	_, __obf_ea722bf866afcda2 := __obf_32414073ae402561[__obf_ce68e709ffad86c7]
	return __obf_ea722bf866afcda2
}

func Run(__obf_ce68e709ffad86c7 string, __obf_aaf970f7207b7551, __obf_9f4c2587953dd0b9 string, __obf_f4c6263732ced88a []string) (string, string, error) {
	return __obf_32414073ae402561[__obf_ce68e709ffad86c7](__obf_aaf970f7207b7551, __obf_9f4c2587953dd0b9, __obf_f4c6263732ced88a)
}
