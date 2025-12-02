package __obf_67ff2b45a2ee7325

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(__obf_56184c3598d8fda1 string) (io.ReadCloser, error) {
	__obf_29c54fad08b62fb3 := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	__obf_1316815cf6b65a5b, __obf_3d6b93bb4b532726 := __obf_29c54fad08b62fb3.Get(__obf_56184c3598d8fda1)
	if __obf_3d6b93bb4b532726 != nil {
		return nil, __obf_3d6b93bb4b532726
	}
	if __obf_1316815cf6b65a5b.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", __obf_1316815cf6b65a5b.StatusCode)
	}
	return __obf_1316815cf6b65a5b.Body, nil
}
