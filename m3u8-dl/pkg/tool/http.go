package __obf_c37d95bc12b416d3

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(__obf_32858afe454e19bd string) (io.ReadCloser, error) {
	__obf_5eb294b3ee00a95c := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	__obf_9082ab3bcbc41225, __obf_3ce485578a57d5a7 := __obf_5eb294b3ee00a95c.Get(__obf_32858afe454e19bd)
	if __obf_3ce485578a57d5a7 != nil {
		return nil, __obf_3ce485578a57d5a7
	}
	if __obf_9082ab3bcbc41225.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", __obf_9082ab3bcbc41225.StatusCode)
	}
	return __obf_9082ab3bcbc41225.Body, nil
}
