package __obf_182a42ae07ff6281

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_b0d252677ecdc968 string, __obf_b821c9b6afd88a60 model.Payload) (__obf_cfccc0cbfb16d21c model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_b821c9b6afd88a60.Files) == 0 {
		__obf_cfccc0cbfb16d21c.
			Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_b821c9b6afd88a60.Lang) {
		__obf_cfccc0cbfb16d21c.
			Error = fmt.Sprintf("lang '%s' is not supported", __obf_b821c9b6afd88a60.Lang)
		return
	}
	__obf_a6051cce2f12ee3a,

		// Write files to disk
		__obf_d91f49d0015e51d4, __obf_1b68fda760a5988a := util.WriteFiles(__obf_b0d252677ecdc968, __obf_b821c9b6afd88a60.Lang+"-", __obf_b821c9b6afd88a60.Files)
	if __obf_1b68fda760a5988a != nil {
		__obf_cfccc0cbfb16d21c.
			Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_1b68fda760a5988a.Error())
		return
	}
	defer os.RemoveAll(__obf_a6051cce2f12ee3a)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_b821c9b6afd88a60.Command == "" {
		__obf_cfccc0cbfb16d21c.
			Stdout, __obf_cfccc0cbfb16d21c.Stderr, __obf_1b68fda760a5988a = lang.Run(__obf_b821c9b6afd88a60.Lang, __obf_a6051cce2f12ee3a, __obf_b821c9b6afd88a60.Stdin, __obf_d91f49d0015e51d4)
	} else {
		__obf_cfccc0cbfb16d21c.
			Stdout, __obf_cfccc0cbfb16d21c.Stderr, __obf_1b68fda760a5988a = cmd.RunBashStdin(__obf_a6051cce2f12ee3a, __obf_b821c9b6afd88a60.Command, __obf_b821c9b6afd88a60.Stdin)
	}
	if __obf_1b68fda760a5988a != nil {
		__obf_cfccc0cbfb16d21c.
			Error = __obf_1b68fda760a5988a.Error()
	}
	return
}
