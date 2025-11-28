package __obf_3a88b426cd9f14e1

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_90206113baabbe94 string, __obf_dd60ec98f3006faf model.Payload) (__obf_23567f7fed847944 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_dd60ec98f3006faf.Files) == 0 {
		__obf_23567f7fed847944.Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_dd60ec98f3006faf.Lang) {
		__obf_23567f7fed847944.Error = fmt.Sprintf("lang '%s' is not supported", __obf_dd60ec98f3006faf.Lang)
		return
	}

	// Write files to disk
	__obf_040e319999e1e623, __obf_5ef5567f74331fb8, __obf_56768e3d0daa4785 := util.WriteFiles(__obf_90206113baabbe94, __obf_dd60ec98f3006faf.Lang+"-", __obf_dd60ec98f3006faf.Files)
	if __obf_56768e3d0daa4785 != nil {
		__obf_23567f7fed847944.Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_56768e3d0daa4785.Error())
		return
	}
	defer os.RemoveAll(__obf_040e319999e1e623)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_dd60ec98f3006faf.Command == "" {
		__obf_23567f7fed847944.Stdout, __obf_23567f7fed847944.Stderr, __obf_56768e3d0daa4785 = lang.Run(__obf_dd60ec98f3006faf.Lang, __obf_040e319999e1e623, __obf_dd60ec98f3006faf.Stdin, __obf_5ef5567f74331fb8)
	} else {
		__obf_23567f7fed847944.Stdout, __obf_23567f7fed847944.Stderr, __obf_56768e3d0daa4785 = cmd.RunBashStdin(__obf_040e319999e1e623, __obf_dd60ec98f3006faf.Command, __obf_dd60ec98f3006faf.Stdin)
	}
	if __obf_56768e3d0daa4785 != nil {
		__obf_23567f7fed847944.Error = __obf_56768e3d0daa4785.Error()
	}
	return
}
