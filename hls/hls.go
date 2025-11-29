package __obf_6ff082bd539c7df0

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
	__obf_b27f3195132067e8 *Worker
	__obf_e76a5f4630004e4b *Worker
	__obf_309e545f7e826626 *Worker
}

func NewHlsServer(__obf_87d5afd76ae831dc *Option) *HlsServer {
	return &HlsServer{__obf_87d5afd76ae831dc, NewWorker(__obf_87d5afd76ae831dc.WorkersNum, __obf_87d5afd76ae831dc.ExeTimeout, __obf_87d5afd76ae831dc.CacheDir),
		NewWorker(__obf_87d5afd76ae831dc.WorkersNum, __obf_87d5afd76ae831dc.ExeTimeout, __obf_87d5afd76ae831dc.CacheDir),
		NewWorker(__obf_87d5afd76ae831dc.WorkersNum, __obf_87d5afd76ae831dc.ExeTimeout, __obf_87d5afd76ae831dc.CacheDir),
	}
}

func ExecJsonStdout(__obf_16a1ed0aadc5fc92 *exec.Cmd, __obf_3082d785fa2087cb any) error {
	__obf_7140f9b71e52894f := new(bytes.Buffer)
	if __obf_ccdf867ff6e9ddb3 := ExecWriteStdout(__obf_16a1ed0aadc5fc92, __obf_7140f9b71e52894f); __obf_ccdf867ff6e9ddb3 != nil {
		return __obf_ccdf867ff6e9ddb3
	}

	return json.Unmarshal(__obf_7140f9b71e52894f.Bytes(), __obf_3082d785fa2087cb)
}

func ExecWriteStdout(__obf_16a1ed0aadc5fc92 *exec.Cmd, __obf_42061c65b3e92c28 io.Writer) error {
	__obf_cb89f3459c117b10, __obf_ccdf867ff6e9ddb3 := __obf_16a1ed0aadc5fc92.StdoutPipe()
	if __obf_ccdf867ff6e9ddb3 != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_ccdf867ff6e9ddb3)
	}
	defer __obf_cb89f3459c117b10.Close()

	if __obf_ccdf867ff6e9ddb3 = __obf_16a1ed0aadc5fc92.Start(); __obf_ccdf867ff6e9ddb3 != nil {
		return fmt.Errorf("error starting command: %v", __obf_ccdf867ff6e9ddb3)
	}

	if _, __obf_ccdf867ff6e9ddb3 = io.Copy(__obf_42061c65b3e92c28, __obf_cb89f3459c117b10); __obf_ccdf867ff6e9ddb3 != nil {
		// Ask the process to exit
		_ = __obf_16a1ed0aadc5fc92.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_16a1ed0aadc5fc92.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_ccdf867ff6e9ddb3)
	}
	if __obf_ccdf867ff6e9ddb3 = __obf_16a1ed0aadc5fc92.Wait(); __obf_ccdf867ff6e9ddb3 != nil {
		return fmt.Errorf("command failed %v", __obf_ccdf867ff6e9ddb3)
	}

	return nil
}
