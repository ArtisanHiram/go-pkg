package __obf_9e2677cd4d2ebe52

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_beedc930bc2c66af string, __obf_9d1e9d10447ed20f model.Payload) (__obf_ae5481cd14005a1c model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_9d1e9d10447ed20f.Files) == 0 {
		__obf_ae5481cd14005a1c.Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_9d1e9d10447ed20f.Lang) {
		__obf_ae5481cd14005a1c.Error = fmt.Sprintf("lang '%s' is not supported", __obf_9d1e9d10447ed20f.Lang)
		return
	}

	// Write files to disk
	__obf_249891adcf5d16f1, __obf_09c55997fe650131, __obf_4ac45fa9bf0dfdf8 := util.WriteFiles(__obf_beedc930bc2c66af, __obf_9d1e9d10447ed20f.Lang+"-", __obf_9d1e9d10447ed20f.Files)
	if __obf_4ac45fa9bf0dfdf8 != nil {
		__obf_ae5481cd14005a1c.Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_4ac45fa9bf0dfdf8.Error())
		return
	}
	defer os.RemoveAll(__obf_249891adcf5d16f1)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_9d1e9d10447ed20f.Command == "" {
		__obf_ae5481cd14005a1c.Stdout, __obf_ae5481cd14005a1c.Stderr, __obf_4ac45fa9bf0dfdf8 = lang.Run(__obf_9d1e9d10447ed20f.Lang, __obf_249891adcf5d16f1, __obf_9d1e9d10447ed20f.Stdin, __obf_09c55997fe650131)
	} else {
		__obf_ae5481cd14005a1c.Stdout, __obf_ae5481cd14005a1c.Stderr, __obf_4ac45fa9bf0dfdf8 = cmd.RunBashStdin(__obf_249891adcf5d16f1, __obf_9d1e9d10447ed20f.Command, __obf_9d1e9d10447ed20f.Stdin)
	}
	if __obf_4ac45fa9bf0dfdf8 != nil {
		__obf_ae5481cd14005a1c.Error = __obf_4ac45fa9bf0dfdf8.Error()
	}
	return
}
