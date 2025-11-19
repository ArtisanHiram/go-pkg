package __obf_dc577bef9ac5c6b8

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
	__obf_b964b68a70df3f8a Cache
	__obf_2b0a27e1d6573003 int
	__obf_770109f9d11f3c66 chan struct{}
}

func NewWorker(__obf_f74f198b56530485, __obf_2b0a27e1d6573003 int, __obf_0047f6edb2231322 string) *Worker {
	__obf_770109f9d11f3c66 := make(chan struct{}, __obf_f74f198b56530485)
	for __obf_5e240e9841a5c237 := __obf_f74f198b56530485; __obf_5e240e9841a5c237 > 0; __obf_5e240e9841a5c237-- {
		__obf_770109f9d11f3c66 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_0047f6edb2231322), __obf_2b0a27e1d6573003, __obf_770109f9d11f3c66}
}

// TODO timeout & context
func (__obf_fdd7931bd5a61fac *Worker) Serve(__obf_adceafa34032b65e string, __obf_2e8b43644037105e []string, __obf_39c0ed11835c4488 io.Writer) error {

	__obf_37493f0bfb931271 := sha1.New()
	__obf_37493f0bfb931271.Write([]byte(__obf_adceafa34032b65e))
	for _, __obf_09a76bbb1dc54f81 := range __obf_2e8b43644037105e {
		__obf_37493f0bfb931271.Write([]byte(__obf_09a76bbb1dc54f81))
	}
	__obf_8a5de64853c61e81 := __obf_37493f0bfb931271.Sum(nil)
	__obf_85503b8843d1609c := fmt.Sprintf("%x", __obf_8a5de64853c61e81)

	__obf_9ba61d882c7140dd, __obf_14a3b54d08bd697d := __obf_fdd7931bd5a61fac.__obf_b964b68a70df3f8a.Get(context.Background(), __obf_85503b8843d1609c)
	// If error getting item, return not served with error
	if __obf_14a3b54d08bd697d != nil {
		return __obf_14a3b54d08bd697d
	}
	// If no item found, return not served with no error
	if __obf_9ba61d882c7140dd != nil {
		// If copying fails, return served with error
		if _, __obf_14a3b54d08bd697d = io.Copy(__obf_39c0ed11835c4488, bytes.NewReader(__obf_9ba61d882c7140dd)); __obf_14a3b54d08bd697d != nil {
			return __obf_14a3b54d08bd697d
		}
		// Everything worked, return served with no error
		return nil
	}

	// Wait for token
	__obf_d1ac06ef8445e5e2 := <-__obf_fdd7931bd5a61fac.__obf_770109f9d11f3c66
	defer func() {
		__obf_fdd7931bd5a61fac.__obf_770109f9d11f3c66 <- __obf_d1ac06ef8445e5e2
	}()

	__obf_76620fff066a9aae := new(bytes.Buffer)
	__obf_48c8e05d734fc616 := io.MultiWriter(__obf_76620fff066a9aae, __obf_39c0ed11835c4488)

	__obf_5d1af4120f069bb9, __obf_22443b10029c5b8b := context.WithTimeout(context.Background(), time.Duration(__obf_fdd7931bd5a61fac.__obf_2b0a27e1d6573003)*time.Second)
	defer __obf_22443b10029c5b8b()

	__obf_e26c9c0f4ab57457 := exec.CommandContext(__obf_5d1af4120f069bb9, __obf_adceafa34032b65e, __obf_2e8b43644037105e...)

	if __obf_14a3b54d08bd697d := ExecWriteStdout(__obf_e26c9c0f4ab57457, __obf_48c8e05d734fc616); __obf_14a3b54d08bd697d != nil {
		return __obf_14a3b54d08bd697d
	}

	return __obf_fdd7931bd5a61fac.__obf_b964b68a70df3f8a.Set(context.Background(), __obf_85503b8843d1609c, __obf_76620fff066a9aae.Bytes())
}
