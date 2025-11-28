package __obf_5028a0a829ddcdab

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
	__obf_68bcf222f9b51eb2 Cache
	__obf_4c670b8ee1e61efe int
	__obf_f1302620d5577eaa chan struct{}
}

func NewWorker(__obf_de29810da2d425cd, __obf_4c670b8ee1e61efe int, __obf_53ae05d42bb5bc2d string) *Worker {
	__obf_f1302620d5577eaa := make(chan struct{}, __obf_de29810da2d425cd)
	for __obf_9ad69c77a03d312e := __obf_de29810da2d425cd; __obf_9ad69c77a03d312e > 0; __obf_9ad69c77a03d312e-- {
		__obf_f1302620d5577eaa <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_53ae05d42bb5bc2d), __obf_4c670b8ee1e61efe, __obf_f1302620d5577eaa}
}

// TODO timeout & context
func (__obf_883d14d3bb9e62d2 *Worker) Serve(__obf_f53647708870c254 string, __obf_4f0e95aaf266f771 []string, __obf_f0018e81daaf1fa9 io.Writer) error {

	__obf_f44570ffde5bd91a := sha1.New()
	__obf_f44570ffde5bd91a.Write([]byte(__obf_f53647708870c254))
	for _, __obf_781de6fadb337177 := range __obf_4f0e95aaf266f771 {
		__obf_f44570ffde5bd91a.Write([]byte(__obf_781de6fadb337177))
	}
	__obf_8910f65fdb98c189 := __obf_f44570ffde5bd91a.Sum(nil)
	__obf_fd503f6151c05aa1 := fmt.Sprintf("%x", __obf_8910f65fdb98c189)

	__obf_83b4446f7bdd5e78, __obf_97be5ef7ba0b1a8f := __obf_883d14d3bb9e62d2.__obf_68bcf222f9b51eb2.Get(context.Background(), __obf_fd503f6151c05aa1)
	// If error getting item, return not served with error
	if __obf_97be5ef7ba0b1a8f != nil {
		return __obf_97be5ef7ba0b1a8f
	}
	// If no item found, return not served with no error
	if __obf_83b4446f7bdd5e78 != nil {
		// If copying fails, return served with error
		if _, __obf_97be5ef7ba0b1a8f = io.Copy(__obf_f0018e81daaf1fa9, bytes.NewReader(__obf_83b4446f7bdd5e78)); __obf_97be5ef7ba0b1a8f != nil {
			return __obf_97be5ef7ba0b1a8f
		}
		// Everything worked, return served with no error
		return nil
	}

	// Wait for token
	__obf_c39e04e68a28b9a0 := <-__obf_883d14d3bb9e62d2.__obf_f1302620d5577eaa
	defer func() {
		__obf_883d14d3bb9e62d2.__obf_f1302620d5577eaa <- __obf_c39e04e68a28b9a0
	}()

	__obf_fb7b061858ce2dfb := new(bytes.Buffer)
	__obf_aeead0a65d36963c := io.MultiWriter(__obf_fb7b061858ce2dfb, __obf_f0018e81daaf1fa9)

	__obf_19b79b6a380e8507, __obf_8c503727c5cd3eea := context.WithTimeout(context.Background(), time.Duration(__obf_883d14d3bb9e62d2.__obf_4c670b8ee1e61efe)*time.Second)
	defer __obf_8c503727c5cd3eea()

	__obf_13daf0dd60267c30 := exec.CommandContext(__obf_19b79b6a380e8507, __obf_f53647708870c254, __obf_4f0e95aaf266f771...)

	if __obf_97be5ef7ba0b1a8f := ExecWriteStdout(__obf_13daf0dd60267c30, __obf_aeead0a65d36963c); __obf_97be5ef7ba0b1a8f != nil {
		return __obf_97be5ef7ba0b1a8f
	}

	return __obf_883d14d3bb9e62d2.__obf_68bcf222f9b51eb2.Set(context.Background(), __obf_fd503f6151c05aa1, __obf_fb7b061858ce2dfb.Bytes())
}
