package __obf_b28324f38df50634

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
	__obf_25e9ffdd6cfa6045 *Worker
	__obf_31c94ec8882b996a *Worker
	__obf_0d5f4cd62e2a5059 *Worker
}

func NewHlsServer(__obf_2a22b5adcfbd0a41 *Option) *HlsServer {
	return &HlsServer{
		__obf_2a22b5adcfbd0a41,
		NewWorker(__obf_2a22b5adcfbd0a41.WorkersNum, __obf_2a22b5adcfbd0a41.ExeTimeout, __obf_2a22b5adcfbd0a41.CacheDir),
		NewWorker(__obf_2a22b5adcfbd0a41.WorkersNum, __obf_2a22b5adcfbd0a41.ExeTimeout, __obf_2a22b5adcfbd0a41.CacheDir),
		NewWorker(__obf_2a22b5adcfbd0a41.WorkersNum, __obf_2a22b5adcfbd0a41.ExeTimeout, __obf_2a22b5adcfbd0a41.CacheDir),
	}
}

func ExecJsonStdout(__obf_5e6b4fe9566ae24a *exec.Cmd, __obf_10deb3d71f9a3d85 any) error {

	__obf_a716135a97d66cdf := new(bytes.Buffer)
	if __obf_a4e073529832bad2 := ExecWriteStdout(__obf_5e6b4fe9566ae24a, __obf_a716135a97d66cdf); __obf_a4e073529832bad2 != nil {
		return __obf_a4e073529832bad2
	}

	return json.Unmarshal(__obf_a716135a97d66cdf.Bytes(), __obf_10deb3d71f9a3d85)
}

func ExecWriteStdout(__obf_5e6b4fe9566ae24a *exec.Cmd, __obf_eb3c523b7bd5ff61 io.Writer) error {

	__obf_e623fc1127e1de26, __obf_a4e073529832bad2 := __obf_5e6b4fe9566ae24a.StdoutPipe()
	if __obf_a4e073529832bad2 != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_a4e073529832bad2)
	}
	defer __obf_e623fc1127e1de26.Close()

	if __obf_a4e073529832bad2 = __obf_5e6b4fe9566ae24a.Start(); __obf_a4e073529832bad2 != nil {
		return fmt.Errorf("error starting command: %v", __obf_a4e073529832bad2)
	}

	if _, __obf_a4e073529832bad2 = io.Copy(__obf_eb3c523b7bd5ff61, __obf_e623fc1127e1de26); __obf_a4e073529832bad2 != nil {
		// Ask the process to exit
		_ = __obf_5e6b4fe9566ae24a.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_5e6b4fe9566ae24a.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_a4e073529832bad2)
	}
	if __obf_a4e073529832bad2 = __obf_5e6b4fe9566ae24a.Wait(); __obf_a4e073529832bad2 != nil {
		return fmt.Errorf("command failed %v", __obf_a4e073529832bad2)
	}

	return nil
}
