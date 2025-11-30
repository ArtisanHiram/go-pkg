package __obf_a0204a07218516a6

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_59a623142268619c string, __obf_5e855f0b81f7edd4 string, __obf_0c9721f091eba392 []string) (string, string, error) {
	__obf_dfbb9c6c41babbe5, __obf_3b8eca5e9fee40d0, __obf_7ce639ea6a22051f := cmd.Run(__obf_59a623142268619c, "javac", __obf_0c9721f091eba392[0])
	if __obf_7ce639ea6a22051f != nil {
		return __obf_dfbb9c6c41babbe5, __obf_3b8eca5e9fee40d0, __obf_7ce639ea6a22051f
	}
	return cmd.RunStdin(__obf_59a623142268619c, __obf_5e855f0b81f7edd4, "java", __obf_8119adeeb55557df(__obf_0c9721f091eba392[0]))
}

func __obf_8119adeeb55557df(__obf_9badb334b93730a9 string) string {
	__obf_9a5c073ff3918f73 := filepath.Ext(__obf_9badb334b93730a9)
	return __obf_9badb334b93730a9[0 : len(__obf_9badb334b93730a9)-len(__obf_9a5c073ff3918f73)]
}
