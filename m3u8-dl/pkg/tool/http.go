package __obf_22f65c3e757f8811

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(__obf_bfc6c4c0e57fa694 string) (io.ReadCloser, error) {
	__obf_1ca4c786916d741b := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	__obf_d3b3c44ab25ee1a4, __obf_f749646f37e10859 := __obf_1ca4c786916d741b.Get(__obf_bfc6c4c0e57fa694)
	if __obf_f749646f37e10859 != nil {
		return nil, __obf_f749646f37e10859
	}
	if __obf_d3b3c44ab25ee1a4.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", __obf_d3b3c44ab25ee1a4.StatusCode)
	}
	return __obf_d3b3c44ab25ee1a4.Body, nil
}
