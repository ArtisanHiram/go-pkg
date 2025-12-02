package __obf_c5c8c536024c8bc3

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_04a14fb0f52d8923 string, __obf_7041769a97a9508a model.Payload) (__obf_2f85dbeee62b20f4 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_7041769a97a9508a.Files) == 0 {
		__obf_2f85dbeee62b20f4.
			Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_7041769a97a9508a.Lang) {
		__obf_2f85dbeee62b20f4.
			Error = fmt.Sprintf("lang '%s' is not supported", __obf_7041769a97a9508a.Lang)
		return
	}
	__obf_ef546632930dd68e,

		// Write files to disk
		__obf_7345682a5f0f4254, __obf_d8621c5be7598851 := util.WriteFiles(__obf_04a14fb0f52d8923, __obf_7041769a97a9508a.Lang+"-", __obf_7041769a97a9508a.Files)
	if __obf_d8621c5be7598851 != nil {
		__obf_2f85dbeee62b20f4.
			Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_d8621c5be7598851.Error())
		return
	}
	defer os.RemoveAll(__obf_ef546632930dd68e)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_7041769a97a9508a.Command == "" {
		__obf_2f85dbeee62b20f4.
			Stdout, __obf_2f85dbeee62b20f4.Stderr, __obf_d8621c5be7598851 = lang.Run(__obf_7041769a97a9508a.Lang, __obf_ef546632930dd68e, __obf_7041769a97a9508a.Stdin, __obf_7345682a5f0f4254)
	} else {
		__obf_2f85dbeee62b20f4.
			Stdout, __obf_2f85dbeee62b20f4.Stderr, __obf_d8621c5be7598851 = cmd.RunBashStdin(__obf_ef546632930dd68e, __obf_7041769a97a9508a.Command, __obf_7041769a97a9508a.Stdin)
	}
	if __obf_d8621c5be7598851 != nil {
		__obf_2f85dbeee62b20f4.
			Error = __obf_d8621c5be7598851.Error()
	}
	return
}
