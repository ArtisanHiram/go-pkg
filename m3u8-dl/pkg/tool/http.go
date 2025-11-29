package __obf_61bc591a775c8985

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(__obf_751af54586959796 string) (io.ReadCloser, error) {
	__obf_61a1a9d200f64ca3 := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	__obf_3e12e47e89812cc0, __obf_48065909968843ec := __obf_61a1a9d200f64ca3.Get(__obf_751af54586959796)
	if __obf_48065909968843ec != nil {
		return nil, __obf_48065909968843ec
	}
	if __obf_3e12e47e89812cc0.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", __obf_3e12e47e89812cc0.StatusCode)
	}
	return __obf_3e12e47e89812cc0.Body, nil
}
