package __obf_aadffefcacf8d80c

import (
	"fmt"
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	lang "github.com/ArtisanHiram/go-pkg/runner/lang"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
	"os"
)

func Execute(__obf_b32f05820c9ef853 string, __obf_1a96f2e2687b6eca model.Payload) (__obf_9db8f11834274464 model.Result) {
	// payload := &model.Payload{}
	// err := json.NewDecoder(os.Stdin).Decode(payload)

	// if err != nil {
	// 	return fmt.Errorf("failed to parse input json (%s)", err.Error())
	// }

	// Ensure that we have at least one file
	if len(__obf_1a96f2e2687b6eca.Files) == 0 {
		__obf_9db8f11834274464.
			Error = "no files given"
		return
	}

	// Check if we support given lang
	if !lang.IsSupported(__obf_1a96f2e2687b6eca.Lang) {
		__obf_9db8f11834274464.
			Error = fmt.Sprintf("lang '%s' is not supported", __obf_1a96f2e2687b6eca.Lang)
		return
	}
	__obf_e8b643c40e8a15f0,

		// Write files to disk
		__obf_a0e5bbff849eed40, __obf_846a02d52baaf6c8 := util.WriteFiles(__obf_b32f05820c9ef853, __obf_1a96f2e2687b6eca.Lang+"-", __obf_1a96f2e2687b6eca.Files)
	if __obf_846a02d52baaf6c8 != nil {
		__obf_9db8f11834274464.
			Error = fmt.Sprintf("failed to write file to disk (%s)", __obf_846a02d52baaf6c8.Error())
		return
	}
	defer os.RemoveAll(__obf_e8b643c40e8a15f0)

	// Execute the given command or run the code with
	// the lang runner if no command is given
	if __obf_1a96f2e2687b6eca.Command == "" {
		__obf_9db8f11834274464.
			Stdout, __obf_9db8f11834274464.Stderr, __obf_846a02d52baaf6c8 = lang.Run(__obf_1a96f2e2687b6eca.Lang, __obf_e8b643c40e8a15f0, __obf_1a96f2e2687b6eca.Stdin, __obf_a0e5bbff849eed40)
	} else {
		__obf_9db8f11834274464.
			Stdout, __obf_9db8f11834274464.Stderr, __obf_846a02d52baaf6c8 = cmd.RunBashStdin(__obf_e8b643c40e8a15f0, __obf_1a96f2e2687b6eca.Command, __obf_1a96f2e2687b6eca.Stdin)
	}
	if __obf_846a02d52baaf6c8 != nil {
		__obf_9db8f11834274464.
			Error = __obf_846a02d52baaf6c8.Error()
	}
	return
}
