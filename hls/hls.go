package __obf_f60326fd90eb13d9

// https://github.com/shimberger/gohls
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"syscall"
)

const (
	FFProbePath = "ffprobe"
	FFMpegPath  = "ffmpeg"
)

type Option struct {
	CacheDir            string  `yaml:"cache-dir"`
	SegmentLen          float32 `yaml:"segment-len"`
	WorkersNum          int     `yaml:"workers-num"`
	ExeTimeout          int     `yaml:"exe-timeout"`
	PlaylistContentType string  `yaml:"playlist-content-type"` //  "application/vnd.apple.mpegurl"
	CaptionContentType  string  `yaml:"caption-content-type"`  //  "text/vtt"
}

type HlsServer struct {
	Option                 *Option
	__obf_e9d8d0cc62a9ece3 *Worker
	__obf_9b2fbfd16e499a01 *Worker
	__obf_69915dd1bdb5a610 *Worker
}

func NewHlsServer(__obf_1730e55a376730d7 *Option) *HlsServer {
	return &HlsServer{__obf_1730e55a376730d7, NewWorker(__obf_1730e55a376730d7.WorkersNum, __obf_1730e55a376730d7.ExeTimeout, __obf_1730e55a376730d7.CacheDir),
		NewWorker(__obf_1730e55a376730d7.WorkersNum, __obf_1730e55a376730d7.ExeTimeout, __obf_1730e55a376730d7.CacheDir),
		NewWorker(__obf_1730e55a376730d7.WorkersNum, __obf_1730e55a376730d7.ExeTimeout, __obf_1730e55a376730d7.CacheDir),
	}
}

func ExecJsonStdout(__obf_faf9d2a7226d34be *exec.Cmd, __obf_db0953cf6d64d018 any) error {
	__obf_a2286bef14c440b9 := new(bytes.Buffer)
	if __obf_083e5685f93ba07c := ExecWriteStdout(__obf_faf9d2a7226d34be, __obf_a2286bef14c440b9); __obf_083e5685f93ba07c != nil {
		return __obf_083e5685f93ba07c
	}

	return json.Unmarshal(__obf_a2286bef14c440b9.Bytes(), __obf_db0953cf6d64d018)
}

func ExecWriteStdout(__obf_faf9d2a7226d34be *exec.Cmd, __obf_811b158c28965ee0 io.Writer) error {
	__obf_88c504c5a516abe0, __obf_083e5685f93ba07c := __obf_faf9d2a7226d34be.StdoutPipe()
	if __obf_083e5685f93ba07c != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_083e5685f93ba07c)
	}
	defer __obf_88c504c5a516abe0.Close()

	if __obf_083e5685f93ba07c = __obf_faf9d2a7226d34be.Start(); __obf_083e5685f93ba07c != nil {
		return fmt.Errorf("error starting command: %v", __obf_083e5685f93ba07c)
	}

	if _, __obf_083e5685f93ba07c = io.Copy(__obf_811b158c28965ee0, __obf_88c504c5a516abe0); __obf_083e5685f93ba07c != nil {
		// Ask the process to exit
		_ = __obf_faf9d2a7226d34be.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_faf9d2a7226d34be.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_083e5685f93ba07c)
	}
	if __obf_083e5685f93ba07c = __obf_faf9d2a7226d34be.Wait(); __obf_083e5685f93ba07c != nil {
		return fmt.Errorf("command failed %v", __obf_083e5685f93ba07c)
	}

	return nil
}
