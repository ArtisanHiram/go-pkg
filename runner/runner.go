package __obf_d36a790d978fa8da

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_1f2e0184a6eba759 string, __obf_6137251eb26a9de9 model.Payload) (__obf_991ef85ced427924 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_6137251eb26a9de9.Files) == 0 {
		__obf_991ef85ced427924.
			Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_6137251eb26a9de9.Lang) {
		__obf_991ef85ced427924.
			Error = fmt.Sprintf("lang '%s' is not supported", __obf_6137251eb26a9de9.Lang)
		return
	}
	__obf_6c41374dfb13ca57,

		// Write files to disk
		__obf_419f746c49088154, __obf_07cdbbe4621d78ff := util.WriteFiles(__obf_1f2e0184a6eba759, __obf_6137251eb26a9de9.Lang+"-", __obf_6137251eb26a9de9.Files)
	if __obf_07cdbbe4621d78ff != nil {
		__obf_991ef85ced427924.
			Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_07cdbbe4621d78ff.Error())
		return
	}
	defer os.RemoveAll(__obf_6c41374dfb13ca57)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_6137251eb26a9de9.Command == "" {
		__obf_991ef85ced427924.
			Stdout, __obf_991ef85ced427924.Stderr, __obf_07cdbbe4621d78ff = lang.Run(__obf_6137251eb26a9de9.Lang, __obf_6c41374dfb13ca57, __obf_6137251eb26a9de9.Stdin, __obf_419f746c49088154)
	} else {
		__obf_991ef85ced427924.
			Stdout, __obf_991ef85ced427924.Stderr, __obf_07cdbbe4621d78ff = cmd.RunBashStdin(__obf_6c41374dfb13ca57, __obf_6137251eb26a9de9.Command, __obf_6137251eb26a9de9.Stdin)
	}
	if __obf_07cdbbe4621d78ff != nil {
		__obf_991ef85ced427924.
			Error = __obf_07cdbbe4621d78ff.Error()
	}
	return
}
