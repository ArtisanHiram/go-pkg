package __obf_dc577bef9ac5c6b8

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
	__obf_b69b54cbf0d766ab *Worker
	__obf_67af47df49797c5b *Worker
	__obf_4def811ff2ac5d05 *Worker
}

func NewHlsServer(__obf_a5243a73d4687ff4 *Option) *HlsServer {
	return &HlsServer{
		__obf_a5243a73d4687ff4,
		NewWorker(__obf_a5243a73d4687ff4.WorkersNum, __obf_a5243a73d4687ff4.ExeTimeout, __obf_a5243a73d4687ff4.CacheDir),
		NewWorker(__obf_a5243a73d4687ff4.WorkersNum, __obf_a5243a73d4687ff4.ExeTimeout, __obf_a5243a73d4687ff4.CacheDir),
		NewWorker(__obf_a5243a73d4687ff4.WorkersNum, __obf_a5243a73d4687ff4.ExeTimeout, __obf_a5243a73d4687ff4.CacheDir),
	}
}

func ExecJsonStdout(__obf_e26c9c0f4ab57457 *exec.Cmd, __obf_09a76bbb1dc54f81 any) error {

	__obf_f5cb3c85b79131f3 := new(bytes.Buffer)
	if __obf_14a3b54d08bd697d := ExecWriteStdout(__obf_e26c9c0f4ab57457, __obf_f5cb3c85b79131f3); __obf_14a3b54d08bd697d != nil {
		return __obf_14a3b54d08bd697d
	}

	return json.Unmarshal(__obf_f5cb3c85b79131f3.Bytes(), __obf_09a76bbb1dc54f81)
}

func ExecWriteStdout(__obf_e26c9c0f4ab57457 *exec.Cmd, __obf_39c0ed11835c4488 io.Writer) error {

	__obf_ef3516da3cb68ce9, __obf_14a3b54d08bd697d := __obf_e26c9c0f4ab57457.StdoutPipe()
	if __obf_14a3b54d08bd697d != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_14a3b54d08bd697d)
	}
	defer __obf_ef3516da3cb68ce9.Close()

	if __obf_14a3b54d08bd697d = __obf_e26c9c0f4ab57457.Start(); __obf_14a3b54d08bd697d != nil {
		return fmt.Errorf("error starting command: %v", __obf_14a3b54d08bd697d)
	}

	if _, __obf_14a3b54d08bd697d = io.Copy(__obf_39c0ed11835c4488, __obf_ef3516da3cb68ce9); __obf_14a3b54d08bd697d != nil {
		// Ask the process to exit
		_ = __obf_e26c9c0f4ab57457.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_e26c9c0f4ab57457.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_14a3b54d08bd697d)
	}
	if __obf_14a3b54d08bd697d = __obf_e26c9c0f4ab57457.Wait(); __obf_14a3b54d08bd697d != nil {
		return fmt.Errorf("command failed %v", __obf_14a3b54d08bd697d)
	}

	return nil
}
