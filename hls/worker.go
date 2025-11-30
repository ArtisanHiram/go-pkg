package __obf_08ee9322ff6adb83

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
	__obf_64cef7953a8944b0 Cache
	__obf_3d775c8dcdb6b5e3 int
	__obf_5e0311aba4f479a3 chan struct{}
}

func NewWorker(__obf_db8c94d23ff312f2, __obf_3d775c8dcdb6b5e3 int, __obf_13389f2f67bd33fb string) *Worker {
	__obf_5e0311aba4f479a3 := make(chan struct{}, __obf_db8c94d23ff312f2)
	for __obf_235c2ebeeb561baf := __obf_db8c94d23ff312f2; __obf_235c2ebeeb561baf > 0; __obf_235c2ebeeb561baf-- {
		__obf_5e0311aba4f479a3 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_13389f2f67bd33fb), __obf_3d775c8dcdb6b5e3,

		// TODO timeout & context
		__obf_5e0311aba4f479a3}
}

func (__obf_504413fedee71524 *Worker) Serve(__obf_ec92df450d34136a string, __obf_307b336feff60aaa []string, __obf_63e18a6991b3d10e io.Writer) error {
	__obf_8414630a3146b4be := sha1.New()
	__obf_8414630a3146b4be.
		Write([]byte(__obf_ec92df450d34136a))
	for _, __obf_53e18a26e170d5aa := range __obf_307b336feff60aaa {
		__obf_8414630a3146b4be.
			Write([]byte(__obf_53e18a26e170d5aa))
	}
	__obf_34162236f11dab8b := __obf_8414630a3146b4be.Sum(nil)
	__obf_f221301302ae8617 := fmt.Sprintf("%x", __obf_34162236f11dab8b)
	__obf_5b3500c3eb0c9d69, __obf_dd2d47498586b1c8 := __obf_504413fedee71524.__obf_64cef7953a8944b0.Get(context.Background(), __obf_f221301302ae8617)
	// If error getting item, return not served with error
	if __obf_dd2d47498586b1c8 != nil {
		return __obf_dd2d47498586b1c8
	}
	// If no item found, return not served with no error
	if __obf_5b3500c3eb0c9d69 != nil {
		// If copying fails, return served with error
		if _, __obf_dd2d47498586b1c8 = io.Copy(__obf_63e18a6991b3d10e, bytes.NewReader(__obf_5b3500c3eb0c9d69)); __obf_dd2d47498586b1c8 != nil {
			return __obf_dd2d47498586b1c8
		}
		// Everything worked, return served with no error
		return nil
	}
	__obf_772d228e741778f1 := // Wait for token
		<-__obf_504413fedee71524.__obf_5e0311aba4f479a3
	defer func() {
		__obf_504413fedee71524.__obf_5e0311aba4f479a3 <- __obf_772d228e741778f1
	}()
	__obf_4bef522c8e46b353 := new(bytes.Buffer)
	__obf_934689fe2f8aeebc := io.MultiWriter(__obf_4bef522c8e46b353, __obf_63e18a6991b3d10e)
	__obf_60ee4500ce182178, __obf_2c533782f8daaa17 := context.WithTimeout(context.Background(), time.Duration(__obf_504413fedee71524.__obf_3d775c8dcdb6b5e3)*time.Second)
	defer __obf_2c533782f8daaa17()
	__obf_e13c1cb8fa57e03f := exec.CommandContext(__obf_60ee4500ce182178, __obf_ec92df450d34136a, __obf_307b336feff60aaa...)

	if __obf_dd2d47498586b1c8 := ExecWriteStdout(__obf_e13c1cb8fa57e03f, __obf_934689fe2f8aeebc); __obf_dd2d47498586b1c8 != nil {
		return __obf_dd2d47498586b1c8
	}

	return __obf_504413fedee71524.__obf_64cef7953a8944b0.Set(context.Background(), __obf_f221301302ae8617, __obf_4bef522c8e46b353.Bytes())
}
