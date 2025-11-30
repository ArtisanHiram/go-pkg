package __obf_f141cf7f40b6e4fc

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_8fd9de122acc3c96 string, __obf_8463053af5bc551e model.Payload) (__obf_545dd3d6543469d2 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_8463053af5bc551e.Files) == 0 {
		__obf_545dd3d6543469d2.
			Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_8463053af5bc551e.Lang) {
		__obf_545dd3d6543469d2.
			Error = fmt.Sprintf("lang '%s' is not supported", __obf_8463053af5bc551e.Lang)
		return
	}
	__obf_629ab314b2a038b8,

		// Write files to disk
		__obf_6cdca50e050611ef, __obf_c9ab8e41f1988438 := util.WriteFiles(__obf_8fd9de122acc3c96, __obf_8463053af5bc551e.Lang+"-", __obf_8463053af5bc551e.Files)
	if __obf_c9ab8e41f1988438 != nil {
		__obf_545dd3d6543469d2.
			Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_c9ab8e41f1988438.Error())
		return
	}
	defer os.RemoveAll(__obf_629ab314b2a038b8)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_8463053af5bc551e.Command == "" {
		__obf_545dd3d6543469d2.
			Stdout, __obf_545dd3d6543469d2.Stderr, __obf_c9ab8e41f1988438 = lang.Run(__obf_8463053af5bc551e.Lang, __obf_629ab314b2a038b8, __obf_8463053af5bc551e.Stdin, __obf_6cdca50e050611ef)
	} else {
		__obf_545dd3d6543469d2.
			Stdout, __obf_545dd3d6543469d2.Stderr, __obf_c9ab8e41f1988438 = cmd.RunBashStdin(__obf_629ab314b2a038b8, __obf_8463053af5bc551e.Command, __obf_8463053af5bc551e.Stdin)
	}
	if __obf_c9ab8e41f1988438 != nil {
		__obf_545dd3d6543469d2.
			Error = __obf_c9ab8e41f1988438.Error()
	}
	return
}
