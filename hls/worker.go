package __obf_f60326fd90eb13d9

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
	__obf_a14bcb6f04fa66e2 Cache
	__obf_80d49490bae8d0f7 int
	__obf_2971d6824cb424a2 chan struct{}
}

func NewWorker(__obf_c0a9a89f1850669c, __obf_80d49490bae8d0f7 int, __obf_10035702a1708104 string) *Worker {
	__obf_2971d6824cb424a2 := make(chan struct{}, __obf_c0a9a89f1850669c)
	for __obf_052c07a19c870a6b := __obf_c0a9a89f1850669c; __obf_052c07a19c870a6b > 0; __obf_052c07a19c870a6b-- {
		__obf_2971d6824cb424a2 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_10035702a1708104), __obf_80d49490bae8d0f7,

		// TODO timeout & context
		__obf_2971d6824cb424a2}
}

func (__obf_dd04e55ec17d0ebf *Worker) Serve(__obf_3d5613944caa05ca string, __obf_11839c275efaf853 []string, __obf_811b158c28965ee0 io.Writer) error {
	__obf_e78d7723f3e23bde := sha1.New()
	__obf_e78d7723f3e23bde.
		Write([]byte(__obf_3d5613944caa05ca))
	for _, __obf_db0953cf6d64d018 := range __obf_11839c275efaf853 {
		__obf_e78d7723f3e23bde.
			Write([]byte(__obf_db0953cf6d64d018))
	}
	__obf_d8815d22cc2245aa := __obf_e78d7723f3e23bde.Sum(nil)
	__obf_6a74cec141e92747 := fmt.Sprintf("%x", __obf_d8815d22cc2245aa)
	__obf_86ddd59adfe77e7d, __obf_083e5685f93ba07c := __obf_dd04e55ec17d0ebf.__obf_a14bcb6f04fa66e2.Get(context.Background(), __obf_6a74cec141e92747)
	// If error getting item, return not served with error
	if __obf_083e5685f93ba07c != nil {
		return __obf_083e5685f93ba07c
	}
	// If no item found, return not served with no error
	if __obf_86ddd59adfe77e7d != nil {
		// If copying fails, return served with error
		if _, __obf_083e5685f93ba07c = io.Copy(__obf_811b158c28965ee0, bytes.NewReader(__obf_86ddd59adfe77e7d)); __obf_083e5685f93ba07c != nil {
			return __obf_083e5685f93ba07c
		}
		// Everything worked, return served with no error
		return nil
	}
	__obf_6cd55942aca56cc0 := // Wait for token
		<-__obf_dd04e55ec17d0ebf.__obf_2971d6824cb424a2
	defer func() {
		__obf_dd04e55ec17d0ebf.__obf_2971d6824cb424a2 <- __obf_6cd55942aca56cc0
	}()
	__obf_e30dadc53473eac4 := new(bytes.Buffer)
	__obf_a54e2e85d03dff69 := io.MultiWriter(__obf_e30dadc53473eac4, __obf_811b158c28965ee0)
	__obf_b580c2427f47a43f, __obf_36b7b94502851f13 := context.WithTimeout(context.Background(), time.Duration(__obf_dd04e55ec17d0ebf.__obf_80d49490bae8d0f7)*time.Second)
	defer __obf_36b7b94502851f13()
	__obf_faf9d2a7226d34be := exec.CommandContext(__obf_b580c2427f47a43f, __obf_3d5613944caa05ca, __obf_11839c275efaf853...)

	if __obf_083e5685f93ba07c := ExecWriteStdout(__obf_faf9d2a7226d34be, __obf_a54e2e85d03dff69); __obf_083e5685f93ba07c != nil {
		return __obf_083e5685f93ba07c
	}

	return __obf_dd04e55ec17d0ebf.__obf_a14bcb6f04fa66e2.Set(context.Background(), __obf_6a74cec141e92747, __obf_e30dadc53473eac4.Bytes())
}
