package __obf_90dd9b56c0f1bd65

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

func (__obf_073dee7a39e9dbe6 *HlsServer) GetSegment(__obf_17a4cba300dc959c string, __obf_f64f4b15a2f386e9, __obf_a20f771403971f41 int, __obf_596cef7dccce398b io.Writer) error {
	__obf_a6b1d1e7d000151b := float32(__obf_f64f4b15a2f386e9) * __obf_073dee7a39e9dbe6.Option.SegmentLen
	__obf_0593c99d6b2a8391 := // see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
		[]string{
			// Prevent encoding to run longer than 30 seonds
			"-timelimit", "45",

			// TODO: Some stuff to investigate
			// "-probesize", "524288",
			// "-fpsprobesize", "10",
			// "-analyzeduration", "2147483647",
			// "-hwaccel:0", "vda",

			// The start time
			// important: needs to be before -i to do input seeking
			"-ss", fmt.Sprintf("%.2f", __obf_a6b1d1e7d000151b),

			// The source file
			"-i", __obf_17a4cba300dc959c, // Put all streams to output
			// "-map", "0",

			// The duration
			"-t", fmt.Sprintf("%.2f", __obf_073dee7a39e9dbe6.Option.SegmentLen),

			// TODO: Find out what it does
			//"-strict", "-2",

			// Synchronize audio
			"-async", "1",

			// 720p
			"-vf", fmt.Sprintf("scale=-2:%d", __obf_a20f771403971f41),

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

			"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_073dee7a39e9dbe6.Option.SegmentLen),

			//"-force_key_frames", "00:00:00.00",
			//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

			//"-f", "mpegts",

			"-f", "ssegment",
			"-segment_time", fmt.Sprintf("%.2f", __obf_073dee7a39e9dbe6.Option.SegmentLen),
			"-initial_offset", fmt.Sprintf("%.2f", __obf_a6b1d1e7d000151b),

			"pipe:out%03d.ts",
		}

	return __obf_073dee7a39e9dbe6.__obf_827fc8ecc572f4c4.Serve(FFMpegPath, __obf_0593c99d6b2a8391, __obf_596cef7dccce398b)
}
