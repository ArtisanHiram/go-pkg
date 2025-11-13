package __obf_ab25ed2437cd567a

import (
	"fmt"
	"io"
)

// var encodeWorker = NewWorker(
// 	2, filepath.Join(".hls", "segments"),
// 	func(i any, w io.Writer) error {
// 		return nil
// 	},
// )

func (__obf_8d3b3856596fc4f4 *HlsServer) GetSegment(__obf_56707f240662d9aa string, __obf_3d86ba35033de9e1, __obf_1b72b8540c667615 int, __obf_98c2f72bee615664 io.Writer) error {
	__obf_0c2392d87add6843 := float32(__obf_3d86ba35033de9e1) * __obf_8d3b3856596fc4f4.Option.SegmentLen
	// see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
	__obf_3a7fd6afab4923e8 := []string{
		// Prevent encoding to run longer than 30 seonds
		"-timelimit", "45",

		// TODO: Some stuff to investigate
		// "-probesize", "524288",
		// "-fpsprobesize", "10",
		// "-analyzeduration", "2147483647",
		// "-hwaccel:0", "vda",

		// The start time
		// important: needs to be before -i to do input seeking
		"-ss", fmt.Sprintf("%.2f", __obf_0c2392d87add6843),

		// The source file
		"-i", __obf_56707f240662d9aa,

		// Put all streams to output
		// "-map", "0",

		// The duration
		"-t", fmt.Sprintf("%.2f", __obf_8d3b3856596fc4f4.Option.SegmentLen),

		// TODO: Find out what it does
		//"-strict", "-2",

		// Synchronize audio
		"-async", "1",

		// 720p
		"-vf", fmt.Sprintf("scale=-2:%d", __obf_1b72b8540c667615),

		// x264 video codec
		"-vcodec", "libx264",

		// x264 preset
		"-preset", "veryfast",

		// aac audio codec
		"-c:a", "aac",
		"-b:a", "128k",
		"-ac", "2",

		// TODO
		"-pix_fmt", "yuv420p",

		//"-r", "25", // fixed framerate

		"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_8d3b3856596fc4f4.Option.SegmentLen),

		//"-force_key_frames", "00:00:00.00",
		//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

		//"-f", "mpegts",

		"-f", "ssegment",
		"-segment_time", fmt.Sprintf("%.2f", __obf_8d3b3856596fc4f4.Option.SegmentLen),
		"-initial_offset", fmt.Sprintf("%.2f", __obf_0c2392d87add6843),

		"pipe:out%03d.ts",
	}

	return __obf_8d3b3856596fc4f4.__obf_f16155fa5237adb9.Serve(FFMpegPath, __obf_3a7fd6afab4923e8, __obf_98c2f72bee615664)
}
