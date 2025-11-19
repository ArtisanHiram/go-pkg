package __obf_fd397140cc63cabf

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_7079de96252aa0e1 string, __obf_d6f2ba1ea6b38b81 model.Payload) (__obf_73e84894b8a22fc4 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_d6f2ba1ea6b38b81.Files) == 0 {
		__obf_73e84894b8a22fc4.Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_d6f2ba1ea6b38b81.Lang) {
		__obf_73e84894b8a22fc4.Error = fmt.Sprintf("lang '%s' is not supported", __obf_d6f2ba1ea6b38b81.Lang)
		return
	}

	// Write files to disk
	__obf_d4a09c5cd06d0ed9, __obf_b5b68727d2470a35, __obf_399a3e633c83a587 := util.WriteFiles(__obf_7079de96252aa0e1, __obf_d6f2ba1ea6b38b81.Lang+"-", __obf_d6f2ba1ea6b38b81.Files)
	if __obf_399a3e633c83a587 != nil {
		__obf_73e84894b8a22fc4.Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_399a3e633c83a587.Error())
		return
	}
	defer os.RemoveAll(__obf_d4a09c5cd06d0ed9)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_d6f2ba1ea6b38b81.Command == "" {
		__obf_73e84894b8a22fc4.Stdout, __obf_73e84894b8a22fc4.Stderr, __obf_399a3e633c83a587 = lang.Run(__obf_d6f2ba1ea6b38b81.Lang, __obf_d4a09c5cd06d0ed9, __obf_d6f2ba1ea6b38b81.Stdin, __obf_b5b68727d2470a35)
	} else {
		__obf_73e84894b8a22fc4.Stdout, __obf_73e84894b8a22fc4.Stderr, __obf_399a3e633c83a587 = cmd.RunBashStdin(__obf_d4a09c5cd06d0ed9, __obf_d6f2ba1ea6b38b81.Command, __obf_d6f2ba1ea6b38b81.Stdin)
	}
	if __obf_399a3e633c83a587 != nil {
		__obf_73e84894b8a22fc4.Error = __obf_399a3e633c83a587.Error()
	}
	return
}
