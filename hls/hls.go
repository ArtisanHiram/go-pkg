package __obf_ab25ed2437cd567a

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
	__obf_4f68dab75aa13315 *Worker
	__obf_57de09c5d2d0d12d *Worker
	__obf_f16155fa5237adb9 *Worker
}

func NewHlsServer(__obf_86bea3f5c5bbf3ed *Option) *HlsServer {
	return &HlsServer{
		__obf_86bea3f5c5bbf3ed,
		NewWorker(__obf_86bea3f5c5bbf3ed.WorkersNum, __obf_86bea3f5c5bbf3ed.ExeTimeout, __obf_86bea3f5c5bbf3ed.CacheDir),
		NewWorker(__obf_86bea3f5c5bbf3ed.WorkersNum, __obf_86bea3f5c5bbf3ed.ExeTimeout, __obf_86bea3f5c5bbf3ed.CacheDir),
		NewWorker(__obf_86bea3f5c5bbf3ed.WorkersNum, __obf_86bea3f5c5bbf3ed.ExeTimeout, __obf_86bea3f5c5bbf3ed.CacheDir),
	}
}

func ExecJsonStdout(__obf_6319c2b6d1f4ed90 *exec.Cmd, __obf_2172876de9966b00 any) error {

	__obf_e097ea2782772156 := new(bytes.Buffer)
	if __obf_9d170493b73636ca := ExecWriteStdout(__obf_6319c2b6d1f4ed90, __obf_e097ea2782772156); __obf_9d170493b73636ca != nil {
		return __obf_9d170493b73636ca
	}

	return json.Unmarshal(__obf_e097ea2782772156.Bytes(), __obf_2172876de9966b00)
}

func ExecWriteStdout(__obf_6319c2b6d1f4ed90 *exec.Cmd, __obf_98c2f72bee615664 io.Writer) error {

	__obf_2fdddcb45a9a7b03, __obf_9d170493b73636ca := __obf_6319c2b6d1f4ed90.StdoutPipe()
	if __obf_9d170493b73636ca != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_9d170493b73636ca)
	}
	defer __obf_2fdddcb45a9a7b03.Close()

	if __obf_9d170493b73636ca = __obf_6319c2b6d1f4ed90.Start(); __obf_9d170493b73636ca != nil {
		return fmt.Errorf("error starting command: %v", __obf_9d170493b73636ca)
	}

	if _, __obf_9d170493b73636ca = io.Copy(__obf_98c2f72bee615664, __obf_2fdddcb45a9a7b03); __obf_9d170493b73636ca != nil {
		// Ask the process to exit
		_ = __obf_6319c2b6d1f4ed90.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_6319c2b6d1f4ed90.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_9d170493b73636ca)
	}
	if __obf_9d170493b73636ca = __obf_6319c2b6d1f4ed90.Wait(); __obf_9d170493b73636ca != nil {
		return fmt.Errorf("command failed %v", __obf_9d170493b73636ca)
	}

	return nil
}
