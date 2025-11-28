package __obf_ffc0a07870d97db2

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_c71c282f9b50a09e string, __obf_90cb55a21c2d3bae model.Payload) (__obf_729126e1762045d6 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_90cb55a21c2d3bae.Files) == 0 {
		__obf_729126e1762045d6.Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_90cb55a21c2d3bae.Lang) {
		__obf_729126e1762045d6.Error = fmt.Sprintf("lang '%s' is not supported", __obf_90cb55a21c2d3bae.Lang)
		return
	}

	// Write files to disk
	__obf_843708c7c305ee9e, __obf_34a1611fbe12752a, __obf_af5e6c2fb4f57e11 := util.WriteFiles(__obf_c71c282f9b50a09e, __obf_90cb55a21c2d3bae.Lang+"-", __obf_90cb55a21c2d3bae.Files)
	if __obf_af5e6c2fb4f57e11 != nil {
		__obf_729126e1762045d6.Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_af5e6c2fb4f57e11.Error())
		return
	}
	defer os.RemoveAll(__obf_843708c7c305ee9e)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_90cb55a21c2d3bae.Command == "" {
		__obf_729126e1762045d6.Stdout, __obf_729126e1762045d6.Stderr, __obf_af5e6c2fb4f57e11 = lang.Run(__obf_90cb55a21c2d3bae.Lang, __obf_843708c7c305ee9e, __obf_90cb55a21c2d3bae.Stdin, __obf_34a1611fbe12752a)
	} else {
		__obf_729126e1762045d6.Stdout, __obf_729126e1762045d6.Stderr, __obf_af5e6c2fb4f57e11 = cmd.RunBashStdin(__obf_843708c7c305ee9e, __obf_90cb55a21c2d3bae.Command, __obf_90cb55a21c2d3bae.Stdin)
	}
	if __obf_af5e6c2fb4f57e11 != nil {
		__obf_729126e1762045d6.Error = __obf_af5e6c2fb4f57e11.Error()
	}
	return
}
