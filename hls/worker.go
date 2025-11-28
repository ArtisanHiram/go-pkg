package __obf_b28324f38df50634

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
	__obf_8539466b2d603433 Cache
	__obf_2aae0e1b9f41e54d int
	__obf_2707183af43d0e2d chan struct{}
}

func NewWorker(__obf_fdb7eb01ed0e7b1a, __obf_2aae0e1b9f41e54d int, __obf_b4ec383c338e5b14 string) *Worker {
	__obf_2707183af43d0e2d := make(chan struct{}, __obf_fdb7eb01ed0e7b1a)
	for __obf_6c11aa123c09e464 := __obf_fdb7eb01ed0e7b1a; __obf_6c11aa123c09e464 > 0; __obf_6c11aa123c09e464-- {
		__obf_2707183af43d0e2d <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_b4ec383c338e5b14), __obf_2aae0e1b9f41e54d, __obf_2707183af43d0e2d}
}

// TODO timeout & context
func (__obf_2dc28af59fafc546 *Worker) Serve(__obf_9311c4af5a096343 string, __obf_95546a22afb69e5e []string, __obf_eb3c523b7bd5ff61 io.Writer) error {

	__obf_1d536cf9c090aa04 := sha1.New()
	__obf_1d536cf9c090aa04.Write([]byte(__obf_9311c4af5a096343))
	for _, __obf_10deb3d71f9a3d85 := range __obf_95546a22afb69e5e {
		__obf_1d536cf9c090aa04.Write([]byte(__obf_10deb3d71f9a3d85))
	}
	__obf_79efae233fbc5315 := __obf_1d536cf9c090aa04.Sum(nil)
	__obf_272d13eb9342c812 := fmt.Sprintf("%x", __obf_79efae233fbc5315)

	__obf_99fca58e1291ee83, __obf_a4e073529832bad2 := __obf_2dc28af59fafc546.__obf_8539466b2d603433.Get(context.Background(), __obf_272d13eb9342c812)
	// If error getting item, return not served with error
	if __obf_a4e073529832bad2 != nil {
		return __obf_a4e073529832bad2
	}
	// If no item found, return not served with no error
	if __obf_99fca58e1291ee83 != nil {
		// If copying fails, return served with error
		if _, __obf_a4e073529832bad2 = io.Copy(__obf_eb3c523b7bd5ff61, bytes.NewReader(__obf_99fca58e1291ee83)); __obf_a4e073529832bad2 != nil {
			return __obf_a4e073529832bad2
		}
		// Everything worked, return served with no error
		return nil
	}

	// Wait for token
	__obf_86086fab27d5b555 := <-__obf_2dc28af59fafc546.__obf_2707183af43d0e2d
	defer func() {
		__obf_2dc28af59fafc546.__obf_2707183af43d0e2d <- __obf_86086fab27d5b555
	}()

	__obf_55571b061f6b7e82 := new(bytes.Buffer)
	__obf_cb4f7201c7accd6c := io.MultiWriter(__obf_55571b061f6b7e82, __obf_eb3c523b7bd5ff61)

	__obf_c2c8a52eb04b2873, __obf_f6957728c129bb63 := context.WithTimeout(context.Background(), time.Duration(__obf_2dc28af59fafc546.__obf_2aae0e1b9f41e54d)*time.Second)
	defer __obf_f6957728c129bb63()

	__obf_5e6b4fe9566ae24a := exec.CommandContext(__obf_c2c8a52eb04b2873, __obf_9311c4af5a096343, __obf_95546a22afb69e5e...)

	if __obf_a4e073529832bad2 := ExecWriteStdout(__obf_5e6b4fe9566ae24a, __obf_cb4f7201c7accd6c); __obf_a4e073529832bad2 != nil {
		return __obf_a4e073529832bad2
	}

	return __obf_2dc28af59fafc546.__obf_8539466b2d603433.Set(context.Background(), __obf_272d13eb9342c812, __obf_55571b061f6b7e82.Bytes())
}
