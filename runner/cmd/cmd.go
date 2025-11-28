package __obf_78b71b10bdae6f5a

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Run(__obf_fb936d2b4b13c9b8 string, __obf_40d09d4e50b87d00 ...string) (string, string, error) {
	return RunStdin(__obf_fb936d2b4b13c9b8, "", __obf_40d09d4e50b87d00...)
}

func RunStdin(__obf_fb936d2b4b13c9b8, __obf_5589752e3b8ec1ba string, __obf_40d09d4e50b87d00 ...string) (string, string, error) {
	// 1. 使用带超时的 Context
	__obf_6d292308a338451f, __obf_a7d28227953f0f76 := context.WithTimeout(context.Background(), 15*time.Second)
	defer __obf_a7d28227953f0f76()

	__obf_78b71b10bdae6f5a := exec.CommandContext(__obf_6d292308a338451f, __obf_40d09d4e50b87d00[0], __obf_40d09d4e50b87d00[1:]...)
	// 2. 捕获输出，避免阻塞
	var __obf_4d26c28b1828aef9, __obf_08fcd25a73dc8c86 bytes.Buffer
	__obf_78b71b10bdae6f5a.Dir = __obf_fb936d2b4b13c9b8
	if __obf_5589752e3b8ec1ba != "" {
		__obf_78b71b10bdae6f5a.Stdin = strings.NewReader(__obf_5589752e3b8ec1ba)
	}
	__obf_78b71b10bdae6f5a.Stdout = &__obf_4d26c28b1828aef9
	__obf_78b71b10bdae6f5a.Stderr = &__obf_08fcd25a73dc8c86

	// 3. 启动并等待
	__obf_fb9dd58905726b26 := __obf_78b71b10bdae6f5a.Run()

	// 4. 如果是超时，Context 会自动杀掉进程
	if __obf_6d292308a338451f.Err() == context.DeadlineExceeded {
		return __obf_4d26c28b1828aef9.String(), __obf_08fcd25a73dc8c86.String(), fmt.Errorf("command timed out")
	}

	return strings.TrimSpace(__obf_4d26c28b1828aef9.String()), __obf_08fcd25a73dc8c86.String(), __obf_fb9dd58905726b26
}

// func RunBash(workDir, command string) (string, string, error) {
// 	return Run(workDir, "bash", "-c", command)
// }

func RunBashStdin(__obf_fb936d2b4b13c9b8, __obf_38c7ced17b878edc, __obf_5589752e3b8ec1ba string) (string, string, error) {
	return RunStdin(__obf_fb936d2b4b13c9b8, __obf_5589752e3b8ec1ba, "bash", "-c", __obf_38c7ced17b878edc)
}
