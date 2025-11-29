package __obf_d46ec55c5ef4935c

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_586a9124022bf13e string, __obf_c83226e738570287 string, __obf_3178da58f50058d3 []string) (string, string, error) {
	__obf_d89fe2958d7240e8 := "main.js"
	__obf_073aec09fa2eb690 := // Find all typescript files and build compile command
		util.FilterByExtension(__obf_3178da58f50058d3, ".ts")
	__obf_9451dca319ecedc6 := append([]string{"tsc", "-out", __obf_d89fe2958d7240e8}, __obf_073aec09fa2eb690...)
	__obf_a9be9bad18257073,

		// Compile to javascript
		__obf_a6c0335e4730c435, __obf_fc1304775289ebff := cmd.Run(__obf_586a9124022bf13e, __obf_9451dca319ecedc6...)
	if __obf_fc1304775289ebff != nil {
		return __obf_a9be9bad18257073, __obf_a6c0335e4730c435, __obf_fc1304775289ebff
	}

	return cmd.RunStdin(__obf_586a9124022bf13e, __obf_c83226e738570287, "node", __obf_d89fe2958d7240e8)
}
