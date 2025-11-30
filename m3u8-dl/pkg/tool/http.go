package __obf_33bc7bf683859fa2

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(__obf_8004495da9eea1f0 string) (io.ReadCloser, error) {
	__obf_5b802ec67c26288f := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	__obf_3747cf33f8360691, __obf_19a0c091a533826e := __obf_5b802ec67c26288f.Get(__obf_8004495da9eea1f0)
	if __obf_19a0c091a533826e != nil {
		return nil, __obf_19a0c091a533826e
	}
	if __obf_3747cf33f8360691.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", __obf_3747cf33f8360691.StatusCode)
	}
	return __obf_3747cf33f8360691.Body, nil
}
