package __obf_90dd9b56c0f1bd65

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
	__obf_b2fb5b3625c77a34 *Worker
	__obf_b62a01073bbe620f *Worker
	__obf_827fc8ecc572f4c4 *Worker
}

func NewHlsServer(__obf_42e9fb7cb040d5ce *Option) *HlsServer {
	return &HlsServer{__obf_42e9fb7cb040d5ce, NewWorker(__obf_42e9fb7cb040d5ce.WorkersNum, __obf_42e9fb7cb040d5ce.ExeTimeout, __obf_42e9fb7cb040d5ce.CacheDir),
		NewWorker(__obf_42e9fb7cb040d5ce.WorkersNum, __obf_42e9fb7cb040d5ce.ExeTimeout, __obf_42e9fb7cb040d5ce.CacheDir),
		NewWorker(__obf_42e9fb7cb040d5ce.WorkersNum, __obf_42e9fb7cb040d5ce.ExeTimeout, __obf_42e9fb7cb040d5ce.CacheDir),
	}
}

func ExecJsonStdout(__obf_ecf73b7cf1ef169f *exec.Cmd, __obf_d6721640180c6d52 any) error {
	__obf_ab9168b9655f4867 := new(bytes.Buffer)
	if __obf_dadfdd29cd0d4fe8 := ExecWriteStdout(__obf_ecf73b7cf1ef169f, __obf_ab9168b9655f4867); __obf_dadfdd29cd0d4fe8 != nil {
		return __obf_dadfdd29cd0d4fe8
	}

	return json.Unmarshal(__obf_ab9168b9655f4867.Bytes(), __obf_d6721640180c6d52)
}

func ExecWriteStdout(__obf_ecf73b7cf1ef169f *exec.Cmd, __obf_596cef7dccce398b io.Writer) error {
	__obf_bc2a30279298a62f, __obf_dadfdd29cd0d4fe8 := __obf_ecf73b7cf1ef169f.StdoutPipe()
	if __obf_dadfdd29cd0d4fe8 != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_dadfdd29cd0d4fe8)
	}
	defer __obf_bc2a30279298a62f.Close()

	if __obf_dadfdd29cd0d4fe8 = __obf_ecf73b7cf1ef169f.Start(); __obf_dadfdd29cd0d4fe8 != nil {
		return fmt.Errorf("error starting command: %v", __obf_dadfdd29cd0d4fe8)
	}

	if _, __obf_dadfdd29cd0d4fe8 = io.Copy(__obf_596cef7dccce398b, __obf_bc2a30279298a62f); __obf_dadfdd29cd0d4fe8 != nil {
		// Ask the process to exit
		_ = __obf_ecf73b7cf1ef169f.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_ecf73b7cf1ef169f.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_dadfdd29cd0d4fe8)
	}
	if __obf_dadfdd29cd0d4fe8 = __obf_ecf73b7cf1ef169f.Wait(); __obf_dadfdd29cd0d4fe8 != nil {
		return fmt.Errorf("command failed %v", __obf_dadfdd29cd0d4fe8)
	}

	return nil
}
