package __obf_b28324f38df50634

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

func (__obf_2dc28af59fafc546 *HlsServer) GetSegment(__obf_a2e9fb1ef0c6fd12 string, __obf_a80863ab8324a6a3, __obf_439b01e73ce35521 int, __obf_eb3c523b7bd5ff61 io.Writer) error {
	__obf_bcbd297e690a29ce := float32(__obf_a80863ab8324a6a3) * __obf_2dc28af59fafc546.Option.SegmentLen
	// see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
	__obf_afb5d9998846f346 := []string{
		// Prevent encoding to run longer than 30 seonds
		"-timelimit", "45",

		// TODO: Some stuff to investigate
		// "-probesize", "524288",
		// "-fpsprobesize", "10",
		// "-analyzeduration", "2147483647",
		// "-hwaccel:0", "vda",

		// The start time
		// important: needs to be before -i to do input seeking
		"-ss", fmt.Sprintf("%.2f", __obf_bcbd297e690a29ce),

		// The source file
		"-i", __obf_a2e9fb1ef0c6fd12,

		// Put all streams to output
		// "-map", "0",

		// The duration
		"-t", fmt.Sprintf("%.2f", __obf_2dc28af59fafc546.Option.SegmentLen),

		// TODO: Find out what it does
		//"-strict", "-2",

		// Synchronize audio
		"-async", "1",

		// 720p
		"-vf", fmt.Sprintf("scale=-2:%d", __obf_439b01e73ce35521),

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

		"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_2dc28af59fafc546.Option.SegmentLen),

		//"-force_key_frames", "00:00:00.00",
		//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

		//"-f", "mpegts",

		"-f", "ssegment",
		"-segment_time", fmt.Sprintf("%.2f", __obf_2dc28af59fafc546.Option.SegmentLen),
		"-initial_offset", fmt.Sprintf("%.2f", __obf_bcbd297e690a29ce),

		"pipe:out%03d.ts",
	}

	return __obf_2dc28af59fafc546.__obf_0d5f4cd62e2a5059.Serve(FFMpegPath, __obf_afb5d9998846f346, __obf_eb3c523b7bd5ff61)
}
