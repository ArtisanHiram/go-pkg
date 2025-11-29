package __obf_6ff082bd539c7df0

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
	__obf_e5fb3085dda36a93 Cache
	__obf_905525ff5fcea8fd int
	__obf_206e0e9a255c14e6 chan struct{}
}

func NewWorker(__obf_b71161f746ac4223, __obf_905525ff5fcea8fd int, __obf_2b482b4cee84b197 string) *Worker {
	__obf_206e0e9a255c14e6 := make(chan struct{}, __obf_b71161f746ac4223)
	for __obf_8b28fed702250682 := __obf_b71161f746ac4223; __obf_8b28fed702250682 > 0; __obf_8b28fed702250682-- {
		__obf_206e0e9a255c14e6 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_2b482b4cee84b197), __obf_905525ff5fcea8fd,

		// TODO timeout & context
		__obf_206e0e9a255c14e6}
}

func (__obf_dd5f757557da2c8e *Worker) Serve(__obf_3724df08f7654bbc string, __obf_7ac5967f5e014a66 []string, __obf_42061c65b3e92c28 io.Writer) error {
	__obf_ebc804b4d166d664 := sha1.New()
	__obf_ebc804b4d166d664.
		Write([]byte(__obf_3724df08f7654bbc))
	for _, __obf_3082d785fa2087cb := range __obf_7ac5967f5e014a66 {
		__obf_ebc804b4d166d664.
			Write([]byte(__obf_3082d785fa2087cb))
	}
	__obf_80ca769be7c85480 := __obf_ebc804b4d166d664.Sum(nil)
	__obf_5efc61a6b0304a66 := fmt.Sprintf("%x", __obf_80ca769be7c85480)
	__obf_d117b00f60680fe5, __obf_ccdf867ff6e9ddb3 := __obf_dd5f757557da2c8e.__obf_e5fb3085dda36a93.Get(context.Background(), __obf_5efc61a6b0304a66)
	// If error getting item, return not served with error
	if __obf_ccdf867ff6e9ddb3 != nil {
		return __obf_ccdf867ff6e9ddb3
	}
	// If no item found, return not served with no error
	if __obf_d117b00f60680fe5 != nil {
		// If copying fails, return served with error
		if _, __obf_ccdf867ff6e9ddb3 = io.Copy(__obf_42061c65b3e92c28, bytes.NewReader(__obf_d117b00f60680fe5)); __obf_ccdf867ff6e9ddb3 != nil {
			return __obf_ccdf867ff6e9ddb3
		}
		// Everything worked, return served with no error
		return nil
	}
	__obf_7fde9fd5516888bc := // Wait for token
		<-__obf_dd5f757557da2c8e.__obf_206e0e9a255c14e6
	defer func() {
		__obf_dd5f757557da2c8e.__obf_206e0e9a255c14e6 <- __obf_7fde9fd5516888bc
	}()
	__obf_2d0af15f23bb4cab := new(bytes.Buffer)
	__obf_bc37dc48bd5e842f := io.MultiWriter(__obf_2d0af15f23bb4cab, __obf_42061c65b3e92c28)
	__obf_5254a2186cef3f5f, __obf_78c0c95a5ece1ca5 := context.WithTimeout(context.Background(), time.Duration(__obf_dd5f757557da2c8e.__obf_905525ff5fcea8fd)*time.Second)
	defer __obf_78c0c95a5ece1ca5()
	__obf_16a1ed0aadc5fc92 := exec.CommandContext(__obf_5254a2186cef3f5f, __obf_3724df08f7654bbc, __obf_7ac5967f5e014a66...)

	if __obf_ccdf867ff6e9ddb3 := ExecWriteStdout(__obf_16a1ed0aadc5fc92, __obf_bc37dc48bd5e842f); __obf_ccdf867ff6e9ddb3 != nil {
		return __obf_ccdf867ff6e9ddb3
	}

	return __obf_dd5f757557da2c8e.__obf_e5fb3085dda36a93.Set(context.Background(), __obf_5efc61a6b0304a66, __obf_2d0af15f23bb4cab.Bytes())
}
