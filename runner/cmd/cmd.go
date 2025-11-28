package __obf_58692bec7796222c

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_b1c853dd80e7728d string, __obf_c92ea10b33006a76 ...string) (string, string, error) {
	return RunStdin(__obf_b1c853dd80e7728d, "", __obf_c92ea10b33006a76...)
}

func RunStdin(__obf_b1c853dd80e7728d, __obf_9ac1582803f5b98a string, __obf_c92ea10b33006a76 ...string) (string, string, error) {
	// 1. 使用带超时的 Context
	__obf_56e80264bc6be04a, __obf_95c8ed15aa6a6cb0 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_95c8ed15aa6a6cb0()

	__obf_58692bec7796222c := exec.CommandContext(__obf_56e80264bc6be04a, __obf_c92ea10b33006a76[0], __obf_c92ea10b33006a76[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_3ec6b496c16c8e88, __obf_26afc774c5a4fa0e bytes.Buffer
	__obf_58692bec7796222c.Dir = __obf_b1c853dd80e7728d
	if __obf_9ac1582803f5b98a != "" {
		__obf_58692bec7796222c.Stdin = strings.NewReader(__obf_9ac1582803f5b98a)
	}
	__obf_58692bec7796222c.Stdout = &__obf_3ec6b496c16c8e88
	__obf_58692bec7796222c.Stderr = &__obf_26afc774c5a4fa0e

	// 3. 启动并等待
	__obf_6e9fd6479b6c0bab := __obf_58692bec7796222c.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_56e80264bc6be04a.Err() == context.DeadlineExceeded {
		return __obf_3ec6b496c16c8e88.String(), __obf_26afc774c5a4fa0e.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_3ec6b496c16c8e88.String()), __obf_26afc774c5a4fa0e.String(), __obf_6e9fd6479b6c0bab
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_b1c853dd80e7728d, __obf_3b3dcc75ccb23639, __obf_9ac1582803f5b98a string) (string, string, error) {
	return RunStdin(__obf_b1c853dd80e7728d, __obf_9ac1582803f5b98a, "bash", "-c", __obf_3b3dcc75ccb23639)
}
