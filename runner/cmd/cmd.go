package __obf_9c4ed0a90dfb1114

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_76f93db30ea31765 string, __obf_1190c99b8ba3f75d ...string) (string, string, error) {
	return RunStdin(__obf_76f93db30ea31765, "", __obf_1190c99b8ba3f75d...)
}

func RunStdin(__obf_76f93db30ea31765, __obf_55272303f56b53e5 string, __obf_1190c99b8ba3f75d ...string) (string, string, error) {
	__obf_b85ddb2aa7250899,
		// 1. 使用带超时的 Context
		__obf_5addb136bdad83c0 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_5addb136bdad83c0()
	__obf_9c4ed0a90dfb1114 := exec.CommandContext(__obf_b85ddb2aa7250899, __obf_1190c99b8ba3f75d[0], __obf_1190c99b8ba3f75d[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_b52a496165c44620, __obf_2aaff9d86c6c2b18 bytes.Buffer
	__obf_9c4ed0a90dfb1114.
		Dir = __obf_76f93db30ea31765
	if __obf_55272303f56b53e5 != "" {
		__obf_9c4ed0a90dfb1114.
			Stdin = strings.NewReader(__obf_55272303f56b53e5)
	}
	__obf_9c4ed0a90dfb1114.
		Stdout = &__obf_b52a496165c44620
	__obf_9c4ed0a90dfb1114.
		Stderr = &__obf_2aaff9d86c6c2b18

	// 3. 启动并等待
	__obf_9075b6abcdcc758d := __obf_9c4ed0a90dfb1114.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_b85ddb2aa7250899.Err() == context.DeadlineExceeded {
		return __obf_b52a496165c44620.String(), __obf_2aaff9d86c6c2b18.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_b52a496165c44620.String()), __obf_2aaff9d86c6c2b18.String(), __obf_9075b6abcdcc758d
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_76f93db30ea31765, __obf_72eac5437885a8f2, __obf_55272303f56b53e5 string) (string, string, error) {
	return RunStdin(__obf_76f93db30ea31765, __obf_55272303f56b53e5, "bash", "-c", __obf_72eac5437885a8f2)
}
