package __obf_ca13a73c9e197bf8

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_ca9d1ecba50d4324 string, __obf_9fad0f65c32699f8 model.Payload) (__obf_c853113e7b0ed655 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_9fad0f65c32699f8.Files) == 0 {
		__obf_c853113e7b0ed655.Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_9fad0f65c32699f8.Lang) {
		__obf_c853113e7b0ed655.Error = fmt.Sprintf("lang '%s' is not supported", __obf_9fad0f65c32699f8.Lang)
		return
	}

	// Write files to disk
	__obf_50fb1629525319d2, __obf_4969ddc4a778cf76, __obf_56c130b13df94410 := util.WriteFiles(__obf_ca9d1ecba50d4324, __obf_9fad0f65c32699f8.Lang+"-", __obf_9fad0f65c32699f8.Files)
	if __obf_56c130b13df94410 != nil {
		__obf_c853113e7b0ed655.Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_56c130b13df94410.Error())
		return
	}
	defer os.RemoveAll(__obf_50fb1629525319d2)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_9fad0f65c32699f8.Command == "" {
		__obf_c853113e7b0ed655.Stdout, __obf_c853113e7b0ed655.Stderr, __obf_56c130b13df94410 = lang.Run(__obf_9fad0f65c32699f8.Lang, __obf_50fb1629525319d2, __obf_9fad0f65c32699f8.Stdin, __obf_4969ddc4a778cf76)
	} else {
		__obf_c853113e7b0ed655.Stdout, __obf_c853113e7b0ed655.Stderr, __obf_56c130b13df94410 = cmd.RunBashStdin(__obf_50fb1629525319d2, __obf_9fad0f65c32699f8.Command, __obf_9fad0f65c32699f8.Stdin)
	}
	if __obf_56c130b13df94410 != nil {
		__obf_c853113e7b0ed655.Error = __obf_56c130b13df94410.Error()
	}
	return
}
