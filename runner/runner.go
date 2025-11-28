package __obf_5d052e2300f267b8

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_2fd9aec7baec7cc6 string, __obf_a84465b4b2dd1160 model.Payload) (__obf_6f8ea2f27f4ef3e3 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_a84465b4b2dd1160.Files) == 0 {
		__obf_6f8ea2f27f4ef3e3.Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_a84465b4b2dd1160.Lang) {
		__obf_6f8ea2f27f4ef3e3.Error = fmt.Sprintf("lang '%s' is not supported", __obf_a84465b4b2dd1160.Lang)
		return
	}

	// Write files to disk
	__obf_740db992e2db7dfe, __obf_75f81bb821c93d50, __obf_7d54196ba6ed28ea := util.WriteFiles(__obf_2fd9aec7baec7cc6, __obf_a84465b4b2dd1160.Lang+"-", __obf_a84465b4b2dd1160.Files)
	if __obf_7d54196ba6ed28ea != nil {
		__obf_6f8ea2f27f4ef3e3.Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_7d54196ba6ed28ea.Error())
		return
	}
	defer os.RemoveAll(__obf_740db992e2db7dfe)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_a84465b4b2dd1160.Command == "" {
		__obf_6f8ea2f27f4ef3e3.Stdout, __obf_6f8ea2f27f4ef3e3.Stderr, __obf_7d54196ba6ed28ea = lang.Run(__obf_a84465b4b2dd1160.Lang, __obf_740db992e2db7dfe, __obf_a84465b4b2dd1160.Stdin, __obf_75f81bb821c93d50)
	} else {
		__obf_6f8ea2f27f4ef3e3.Stdout, __obf_6f8ea2f27f4ef3e3.Stderr, __obf_7d54196ba6ed28ea = cmd.RunBashStdin(__obf_740db992e2db7dfe, __obf_a84465b4b2dd1160.Command, __obf_a84465b4b2dd1160.Stdin)
	}
	if __obf_7d54196ba6ed28ea != nil {
		__obf_6f8ea2f27f4ef3e3.Error = __obf_7d54196ba6ed28ea.Error()
	}
	return
}
