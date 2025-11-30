package __obf_90dd9b56c0f1bd65

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
	__obf_2a4364df72614e13 Cache
	__obf_40a6146487dfd586 int
	__obf_db74510fb1ed2add chan struct{}
}

func NewWorker(__obf_410161e5b0d47e0e, __obf_40a6146487dfd586 int, __obf_45281391728ee624 string) *Worker {
	__obf_db74510fb1ed2add := make(chan struct{}, __obf_410161e5b0d47e0e)
	for __obf_846eca422aa6eb22 := __obf_410161e5b0d47e0e; __obf_846eca422aa6eb22 > 0; __obf_846eca422aa6eb22-- {
		__obf_db74510fb1ed2add <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_45281391728ee624), __obf_40a6146487dfd586,

		// TODO timeout & context
		__obf_db74510fb1ed2add}
}

func (__obf_073dee7a39e9dbe6 *Worker) Serve(__obf_23630efb813f591b string, __obf_54c5c89117e7cf86 []string, __obf_596cef7dccce398b io.Writer) error {
	__obf_f1aa14eb0b0dc249 := sha1.New()
	__obf_f1aa14eb0b0dc249.
		Write([]byte(__obf_23630efb813f591b))
	for _, __obf_d6721640180c6d52 := range __obf_54c5c89117e7cf86 {
		__obf_f1aa14eb0b0dc249.
			Write([]byte(__obf_d6721640180c6d52))
	}
	__obf_8fb1b0583f96edc7 := __obf_f1aa14eb0b0dc249.Sum(nil)
	__obf_592914ec2d7c3fe8 := fmt.Sprintf("%x", __obf_8fb1b0583f96edc7)
	__obf_d04ccaf0c7dc23d0, __obf_dadfdd29cd0d4fe8 := __obf_073dee7a39e9dbe6.__obf_2a4364df72614e13.Get(context.Background(), __obf_592914ec2d7c3fe8)
	// If error getting item, return not served with error
	if __obf_dadfdd29cd0d4fe8 != nil {
		return __obf_dadfdd29cd0d4fe8
	}
	// If no item found, return not served with no error
	if __obf_d04ccaf0c7dc23d0 != nil {
		// If copying fails, return served with error
		if _, __obf_dadfdd29cd0d4fe8 = io.Copy(__obf_596cef7dccce398b, bytes.NewReader(__obf_d04ccaf0c7dc23d0)); __obf_dadfdd29cd0d4fe8 != nil {
			return __obf_dadfdd29cd0d4fe8
		}
		// Everything worked, return served with no error
		return nil
	}
	__obf_38644ab655175231 := // Wait for token
		<-__obf_073dee7a39e9dbe6.__obf_db74510fb1ed2add
	defer func() {
		__obf_073dee7a39e9dbe6.__obf_db74510fb1ed2add <- __obf_38644ab655175231
	}()
	__obf_34cdc4a884f7d3e6 := new(bytes.Buffer)
	__obf_7cbd54c06caa7766 := io.MultiWriter(__obf_34cdc4a884f7d3e6, __obf_596cef7dccce398b)
	__obf_9b5717753d5f5860, __obf_17c78879ddf4acdb := context.WithTimeout(context.Background(), time.Duration(__obf_073dee7a39e9dbe6.__obf_40a6146487dfd586)*time.Second)
	defer __obf_17c78879ddf4acdb()
	__obf_ecf73b7cf1ef169f := exec.CommandContext(__obf_9b5717753d5f5860, __obf_23630efb813f591b, __obf_54c5c89117e7cf86...)

	if __obf_dadfdd29cd0d4fe8 := ExecWriteStdout(__obf_ecf73b7cf1ef169f, __obf_7cbd54c06caa7766); __obf_dadfdd29cd0d4fe8 != nil {
		return __obf_dadfdd29cd0d4fe8
	}

	return __obf_073dee7a39e9dbe6.__obf_2a4364df72614e13.Set(context.Background(), __obf_592914ec2d7c3fe8, __obf_34cdc4a884f7d3e6.Bytes())
}
