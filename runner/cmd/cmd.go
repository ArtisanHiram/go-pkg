package __obf_16c714e8f873bafb

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_2ba3bec82c8c4cfb string, __obf_d6ae5347d1c8275d ...string) (string, string, error) {
	return RunStdin(__obf_2ba3bec82c8c4cfb, "", __obf_d6ae5347d1c8275d...)
}

func RunStdin(__obf_2ba3bec82c8c4cfb, __obf_164ecd02a6001b02 string, __obf_d6ae5347d1c8275d ...string) (string, string, error) {
	// 1. 使用带超时的 Context
	__obf_43d69dfc747a2088, __obf_eff7460f7dd9a02d := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_eff7460f7dd9a02d()

	__obf_16c714e8f873bafb := exec.CommandContext(__obf_43d69dfc747a2088, __obf_d6ae5347d1c8275d[0], __obf_d6ae5347d1c8275d[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_ecc3e787a038a09a, __obf_97b31e8f18747d0d bytes.Buffer
	__obf_16c714e8f873bafb.Dir = __obf_2ba3bec82c8c4cfb
	if __obf_164ecd02a6001b02 != "" {
		__obf_16c714e8f873bafb.Stdin = strings.NewReader(__obf_164ecd02a6001b02)
	}
	__obf_16c714e8f873bafb.Stdout = &__obf_ecc3e787a038a09a
	__obf_16c714e8f873bafb.Stderr = &__obf_97b31e8f18747d0d

	// 3. 启动并等待
	__obf_a46e1877d8e657f4 := __obf_16c714e8f873bafb.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_43d69dfc747a2088.Err() == context.DeadlineExceeded {
		return __obf_ecc3e787a038a09a.String(), __obf_97b31e8f18747d0d.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_ecc3e787a038a09a.String()), __obf_97b31e8f18747d0d.String(), __obf_a46e1877d8e657f4
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_2ba3bec82c8c4cfb, __obf_4b3ea61172c576ca, __obf_164ecd02a6001b02 string) (string, string, error) {
	return RunStdin(__obf_2ba3bec82c8c4cfb, __obf_164ecd02a6001b02, "bash", "-c", __obf_4b3ea61172c576ca)
}
