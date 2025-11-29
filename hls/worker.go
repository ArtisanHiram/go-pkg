package __obf_5441fcd9a319cf59

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
	__obf_f7796e5ef5512d18 Cache
	__obf_2a2f3c55da4e482b int
	__obf_82a2658457321852 chan struct{}
}

func NewWorker(__obf_3a824ddc0e6137ce, __obf_2a2f3c55da4e482b int, __obf_62f1dea21133ae23 string) *Worker {
	__obf_82a2658457321852 := make(chan struct{}, __obf_3a824ddc0e6137ce)
	for __obf_74b902588e070f86 := __obf_3a824ddc0e6137ce; __obf_74b902588e070f86 > 0; __obf_74b902588e070f86-- {
		__obf_82a2658457321852 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_62f1dea21133ae23), __obf_2a2f3c55da4e482b,

		// TODO timeout & context
		__obf_82a2658457321852}
}

func (__obf_bf443cef12bfef60 *Worker) Serve(__obf_0359b3d003ec02e1 string, __obf_76b3ec2f9f2cf70c []string, __obf_9e22338370faf798 io.Writer) error {
	__obf_da2419ec6684d5ee := sha1.New()
	__obf_da2419ec6684d5ee.
		Write([]byte(__obf_0359b3d003ec02e1))
	for _, __obf_f86135f935731f8f := range __obf_76b3ec2f9f2cf70c {
		__obf_da2419ec6684d5ee.
			Write([]byte(__obf_f86135f935731f8f))
	}
	__obf_ae78b96c34cc482c := __obf_da2419ec6684d5ee.Sum(nil)
	__obf_d7f29c2f26668a2e := fmt.Sprintf("%x", __obf_ae78b96c34cc482c)
	__obf_b66b1e1568d113aa, __obf_e8125dea727cd4c9 := __obf_bf443cef12bfef60.__obf_f7796e5ef5512d18.Get(context.Background(), __obf_d7f29c2f26668a2e)
	// If error getting item, return not served with error
	if __obf_e8125dea727cd4c9 != nil {
		return __obf_e8125dea727cd4c9
	}
	// If no item found, return not served with no error
	if __obf_b66b1e1568d113aa != nil {
		// If copying fails, return served with error
		if _, __obf_e8125dea727cd4c9 = io.Copy(__obf_9e22338370faf798, bytes.NewReader(__obf_b66b1e1568d113aa)); __obf_e8125dea727cd4c9 != nil {
			return __obf_e8125dea727cd4c9
		}
		// Everything worked, return served with no error
		return nil
	}
	__obf_8082e60f7ad9329f := // Wait for token
		<-__obf_bf443cef12bfef60.__obf_82a2658457321852
	defer func() {
		__obf_bf443cef12bfef60.__obf_82a2658457321852 <- __obf_8082e60f7ad9329f
	}()
	__obf_f9903797af6a6b86 := new(bytes.Buffer)
	__obf_b79858cb88c5bf48 := io.MultiWriter(__obf_f9903797af6a6b86, __obf_9e22338370faf798)
	__obf_83fad17274e83f8e, __obf_79bf09d29303d0f4 := context.WithTimeout(context.Background(), time.Duration(__obf_bf443cef12bfef60.__obf_2a2f3c55da4e482b)*time.Second)
	defer __obf_79bf09d29303d0f4()
	__obf_85b3eedfe1a05aac := exec.CommandContext(__obf_83fad17274e83f8e, __obf_0359b3d003ec02e1, __obf_76b3ec2f9f2cf70c...)

	if __obf_e8125dea727cd4c9 := ExecWriteStdout(__obf_85b3eedfe1a05aac, __obf_b79858cb88c5bf48); __obf_e8125dea727cd4c9 != nil {
		return __obf_e8125dea727cd4c9
	}

	return __obf_bf443cef12bfef60.__obf_f7796e5ef5512d18.Set(context.Background(), __obf_d7f29c2f26668a2e, __obf_f9903797af6a6b86.Bytes())
}
