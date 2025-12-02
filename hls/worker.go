package __obf_eb5e805b9fc230e3

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"os/exec"
	"time"
)

// type Handler func(any, io.Writer) error

type Worker struct {
	__obf_ee96e3a21e7a3d1e Cache
	__obf_dcc34cf142bcd8e3 int
	__obf_56dd4800fe78e956 chan struct{}
}

func NewWorker(__obf_969cb0067c7534fa, __obf_dcc34cf142bcd8e3 int, __obf_e696f94162cbd4e1 string) *Worker {
	__obf_56dd4800fe78e956 := make(chan struct{}, __obf_969cb0067c7534fa)
	for __obf_066c50b91b34eb64 := __obf_969cb0067c7534fa; __obf_066c50b91b34eb64 > 0; __obf_066c50b91b34eb64-- {
		__obf_56dd4800fe78e956 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_e696f94162cbd4e1), __obf_dcc34cf142bcd8e3,

		// TODO timeout & context
		__obf_56dd4800fe78e956}
}

func (__obf_f3546e71ce2e1d63 *Worker) Serve(__obf_fe782748303d0618 string, __obf_0fe133ec36d3a08e []string, __obf_1cab0e86de8129e5 io.Writer) error {
	__obf_0c996dbdc32c6525 := sha1.New()
	__obf_0c996dbdc32c6525.
		Write([]byte(__obf_fe782748303d0618))
	for _, __obf_36f9f90fc6663db6 := range __obf_0fe133ec36d3a08e {
		__obf_0c996dbdc32c6525.
			Write([]byte(__obf_36f9f90fc6663db6))
	}
	__obf_a95b2fd64a87e776 := __obf_0c996dbdc32c6525.Sum(nil)
	__obf_2b7157422f48db91 := fmt.Sprintf("%x", __obf_a95b2fd64a87e776)
	__obf_6cab1b7f43fff094, __obf_2d43222a56faebee := __obf_f3546e71ce2e1d63.__obf_ee96e3a21e7a3d1e.Get(context.Background(), __obf_2b7157422f48db91)
	// If error getting item, return not served with error
	if __obf_2d43222a56faebee != nil {
		return __obf_2d43222a56faebee
	}
	// If no item found, return not served with no error
	if __obf_6cab1b7f43fff094 != nil {
		// If copying fails, return served with error
		if _, __obf_2d43222a56faebee = io.Copy(__obf_1cab0e86de8129e5, bytes.NewReader(__obf_6cab1b7f43fff094)); __obf_2d43222a56faebee != nil {
			return __obf_2d43222a56faebee
		}
		// Everything worked, return served with no error
		return nil
	}
	__obf_6bfaffe2b36eacbe := // Wait for token
		<-__obf_f3546e71ce2e1d63.__obf_56dd4800fe78e956
	defer func() {
		__obf_f3546e71ce2e1d63.__obf_56dd4800fe78e956 <- __obf_6bfaffe2b36eacbe
	}()
	__obf_a74b415b17ac0654 := new(bytes.Buffer)
	__obf_86c941119a906475 := io.MultiWriter(__obf_a74b415b17ac0654, __obf_1cab0e86de8129e5)
	__obf_5f0e28df6b1ffd43, __obf_ac08eb4ed2a9a55f := context.WithTimeout(context.Background(), time.Duration(__obf_f3546e71ce2e1d63.__obf_dcc34cf142bcd8e3)*time.Second)
	defer __obf_ac08eb4ed2a9a55f()
	__obf_81bd5123537bea4c := exec.CommandContext(__obf_5f0e28df6b1ffd43, __obf_fe782748303d0618, __obf_0fe133ec36d3a08e...)

	if __obf_2d43222a56faebee := ExecWriteStdout(__obf_81bd5123537bea4c, __obf_86c941119a906475); __obf_2d43222a56faebee != nil {
		return __obf_2d43222a56faebee
	}

	return __obf_f3546e71ce2e1d63.__obf_ee96e3a21e7a3d1e.Set(context.Background(), __obf_2b7157422f48db91, __obf_a74b415b17ac0654.Bytes())
}
