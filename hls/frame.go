package __obf_5441fcd9a319cf59

import (
	"fmt"
	"io"
)

func (__obf_bf443cef12bfef60 *HlsServer) GetFrame(__obf_7259e223b27c51cd string, __obf_f11c8a13b0a7559d int, __obf_9e22338370faf798 io.Writer) error {
	__obf_e5483301bac834be := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_f11c8a13b0a7559d),
		"-i", __obf_7259e223b27c51cd, "-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_bf443cef12bfef60.__obf_73e20fe74a76e051.Serve(FFMpegPath, __obf_e5483301bac834be, __obf_9e22338370faf798)
}
