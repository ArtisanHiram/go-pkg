package __obf_f52044515183674f

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_f6f04d53e8ab8905 string, __obf_ed2b863761d4a389 model.Payload) (__obf_8e97fc13017e91ea model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_ed2b863761d4a389.Files) == 0 {
		__obf_8e97fc13017e91ea.
			Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_ed2b863761d4a389.Lang) {
		__obf_8e97fc13017e91ea.
			Error = fmt.Sprintf("lang '%s' is not supported", __obf_ed2b863761d4a389.Lang)
		return
	}
	__obf_1bdc9f4455b83b4d,

		// Write files to disk
		__obf_8ba45b71e26a5455, __obf_fd0ab35b33b997f0 := util.WriteFiles(__obf_f6f04d53e8ab8905, __obf_ed2b863761d4a389.Lang+"-", __obf_ed2b863761d4a389.Files)
	if __obf_fd0ab35b33b997f0 != nil {
		__obf_8e97fc13017e91ea.
			Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_fd0ab35b33b997f0.Error())
		return
	}
	defer os.RemoveAll(__obf_1bdc9f4455b83b4d)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_ed2b863761d4a389.Command == "" {
		__obf_8e97fc13017e91ea.
			Stdout, __obf_8e97fc13017e91ea.Stderr, __obf_fd0ab35b33b997f0 = lang.Run(__obf_ed2b863761d4a389.Lang, __obf_1bdc9f4455b83b4d, __obf_ed2b863761d4a389.Stdin, __obf_8ba45b71e26a5455)
	} else {
		__obf_8e97fc13017e91ea.
			Stdout, __obf_8e97fc13017e91ea.Stderr, __obf_fd0ab35b33b997f0 = cmd.RunBashStdin(__obf_1bdc9f4455b83b4d, __obf_ed2b863761d4a389.Command, __obf_ed2b863761d4a389.Stdin)
	}
	if __obf_fd0ab35b33b997f0 != nil {
		__obf_8e97fc13017e91ea.
			Error = __obf_fd0ab35b33b997f0.Error()
	}
	return
}
