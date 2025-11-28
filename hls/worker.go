package __obf_acea4ab24a824c18

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
	__obf_fa45b4ca2954ce06 Cache
	__obf_73b090ce768636ed int
	__obf_48f29e08d7556e48 chan struct{}
}

func NewWorker(__obf_c22732ae4ecf7ee8, __obf_73b090ce768636ed int, __obf_f61e6cab3af72bd7 string) *Worker {
	__obf_48f29e08d7556e48 := make(chan struct{}, __obf_c22732ae4ecf7ee8)
	for __obf_2ea2f1bc4bda8683 := __obf_c22732ae4ecf7ee8; __obf_2ea2f1bc4bda8683 > 0; __obf_2ea2f1bc4bda8683-- {
		__obf_48f29e08d7556e48 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_f61e6cab3af72bd7), __obf_73b090ce768636ed, __obf_48f29e08d7556e48}
}

// TODO timeout & context
func (__obf_c971f1d170edee24 *Worker) Serve(__obf_ee6ae499b69be927 string, __obf_a3b3d2350c0e799f []string, __obf_7bfc44ad58c48031 io.Writer) error {

	__obf_7c2178dac28566c4 := sha1.New()
	__obf_7c2178dac28566c4.Write([]byte(__obf_ee6ae499b69be927))
	for _, __obf_3699794f57b2a656 := range __obf_a3b3d2350c0e799f {
		__obf_7c2178dac28566c4.Write([]byte(__obf_3699794f57b2a656))
	}
	__obf_2859499d7b3fbdb2 := __obf_7c2178dac28566c4.Sum(nil)
	__obf_f1b61571a6fd5b78 := fmt.Sprintf("%x", __obf_2859499d7b3fbdb2)

	__obf_09d4c86b0d837031, __obf_3698cbf06506c070 := __obf_c971f1d170edee24.__obf_fa45b4ca2954ce06.Get(context.Background(), __obf_f1b61571a6fd5b78)
	// If error getting item, return not served with error
	if __obf_3698cbf06506c070 != nil {
		return __obf_3698cbf06506c070
	}
	// If no item found, return not served with no error
	if __obf_09d4c86b0d837031 != nil {
		// If copying fails, return served with error
		if _, __obf_3698cbf06506c070 = io.Copy(__obf_7bfc44ad58c48031, bytes.NewReader(__obf_09d4c86b0d837031)); __obf_3698cbf06506c070 != nil {
			return __obf_3698cbf06506c070
		}
		// Everything worked, return served with no error
		return nil
	}

	// Wait for token
	__obf_89e675a576732b01 := <-__obf_c971f1d170edee24.__obf_48f29e08d7556e48
	defer func() {
		__obf_c971f1d170edee24.__obf_48f29e08d7556e48 <- __obf_89e675a576732b01
	}()

	__obf_e49addece5bb626d := new(bytes.Buffer)
	__obf_2b71608a7ec70549 := io.MultiWriter(__obf_e49addece5bb626d, __obf_7bfc44ad58c48031)

	__obf_efbbc02bedef0e2d, __obf_0e915d14e5e56b99 := context.WithTimeout(context.Background(), time.Duration(__obf_c971f1d170edee24.__obf_73b090ce768636ed)*time.Second)
	defer __obf_0e915d14e5e56b99()

	__obf_d72bd3c5aff9de25 := exec.CommandContext(__obf_efbbc02bedef0e2d, __obf_ee6ae499b69be927, __obf_a3b3d2350c0e799f...)

	if __obf_3698cbf06506c070 := ExecWriteStdout(__obf_d72bd3c5aff9de25, __obf_2b71608a7ec70549); __obf_3698cbf06506c070 != nil {
		return __obf_3698cbf06506c070
	}

	return __obf_c971f1d170edee24.__obf_fa45b4ca2954ce06.Set(context.Background(), __obf_f1b61571a6fd5b78, __obf_e49addece5bb626d.Bytes())
}
