package __obf_6f9b75fc21ef8ef6

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(__obf_2213dcf7603e240c string) (io.ReadCloser, error) {
	__obf_5e64f9c11624866e := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	__obf_6a7b6dc310450e5d, __obf_8eefb2cd18f480f3 := __obf_5e64f9c11624866e.Get(__obf_2213dcf7603e240c)
	if __obf_8eefb2cd18f480f3 != nil {
		return nil, __obf_8eefb2cd18f480f3
	}
	if __obf_6a7b6dc310450e5d.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", __obf_6a7b6dc310450e5d.StatusCode)
	}
	return __obf_6a7b6dc310450e5d.Body, nil
}
