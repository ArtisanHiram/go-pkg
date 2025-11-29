package __obf_b9b152515ee15224

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_10ccb3ba37424578 string, __obf_2953c0cee9d207dd ...string) (string, string, error) {
	return RunStdin(__obf_10ccb3ba37424578, "", __obf_2953c0cee9d207dd...)
}

func RunStdin(__obf_10ccb3ba37424578, __obf_666cf5f2b0e0f698 string, __obf_2953c0cee9d207dd ...string) (string, string, error) {
	__obf_83d1227dde133840,
		// 1. 使用带超时的 Context
		__obf_9a31d622aca4f5b0 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_9a31d622aca4f5b0()
	__obf_b9b152515ee15224 := exec.CommandContext(__obf_83d1227dde133840, __obf_2953c0cee9d207dd[0], __obf_2953c0cee9d207dd[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_5ba763d45e7f0f30, __obf_95d388ae68ce5ee6 bytes.Buffer
	__obf_b9b152515ee15224.
		Dir = __obf_10ccb3ba37424578
	if __obf_666cf5f2b0e0f698 != "" {
		__obf_b9b152515ee15224.
			Stdin = strings.NewReader(__obf_666cf5f2b0e0f698)
	}
	__obf_b9b152515ee15224.
		Stdout = &__obf_5ba763d45e7f0f30
	__obf_b9b152515ee15224.
		Stderr = &__obf_95d388ae68ce5ee6

	// 3. 启动并等待
	__obf_95a2cf0030fd430d := __obf_b9b152515ee15224.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_83d1227dde133840.Err() == context.DeadlineExceeded {
		return __obf_5ba763d45e7f0f30.String(), __obf_95d388ae68ce5ee6.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_5ba763d45e7f0f30.String()), __obf_95d388ae68ce5ee6.String(), __obf_95a2cf0030fd430d
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_10ccb3ba37424578, __obf_0b3384f8a6cad6ab, __obf_666cf5f2b0e0f698 string) (string, string, error) {
	return RunStdin(__obf_10ccb3ba37424578, __obf_666cf5f2b0e0f698, "bash", "-c", __obf_0b3384f8a6cad6ab)
}
